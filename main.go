package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/khatibomar/chunky/check"
	"github.com/khatibomar/chunky/dwn"
)

var (
	ErrMissingArgs = errors.New("Missing or Invalid arguments")
)

const (
	description = `
chunky is a tool that will allow you to download subscribers only videos from twitch.
The tool is under current developement so many bugs will occur , and many missing features and many hard coded stuff.
So please if you found any bug or missing features feel free to open an issue on project page:
https://github.com/khatibomar/chunky
`
)

func main() {
	link := flag.String("url", "", `provide a link that have a chunk , example:
https://d2nvs31859zcd8.cloudfront.net/70c102b5b66dbeac89e4_channel_name_blaabllablablabl/chunked/X.ts
`)
	p := flag.String("dir", "", `specify a download path , for *nix users use $HOME instead of ~
In case no absolute path specified the folder will be created in same dir as the tool folder`)
	name := flag.String("name", "", "the name you want to save the video with without .mp4")
	max := flag.Int("max", -1, "provide the excpected max number of files, zero or negative numbers will be treated as max int")
	down := flag.Bool("dwn", true, "put -down=false if you just want to get chunks size without downloading files")

	flag.Parse()

	*p = path.Clean(*p)

	c := Config{
		Link: *link,
		Max:  *max,
		Path: *p,
		Dwn:  *down,
		Name: *name,
	}

	infoLogger := log.New(os.Stdout, "Info : ", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "Error : ", log.Ldate|log.Ltime)
	errChan := make(chan error)
	doneChan := make(chan bool)
	chunksChan := make(chan int, 1)

	if *down {
		if err := checkDependencies(); err != nil {
			errLogger.Println(err)
			errLogger.Fatal("It look like ffmpeg is not install it, please installed before using the app")
		}
	}

	go run(infoLogger, c, chunksChan, errChan, doneChan)

	for {
		select {
		case err := <-errChan:
			if errors.Is(err, ErrMissingArgs) {
				fmt.Println(description)
				fmt.Println("Usage: ")
				flag.PrintDefaults()
				return
			}
			if err != nil {
				if errors.Is(errors.Unwrap(err), dwn.ErrFileExist) {
					infoLogger.Println(err)
				} else {
					errLogger.Fatal(err)
				}
			}
		case done := <-doneChan:
			if done {
				close(errChan)
				close(doneChan)
				goto DONE_FOR
			}
		}
	}
DONE_FOR:
	if *down {
		ffmpegPath, err := exec.LookPath("ffmpeg")
		if err != nil {
			errLogger.Fatalln("ffmpeg is not installed , please install it")
		}
		chunksNb := <-chunksChan
		close(chunksChan)
		f, err := ioutil.TempFile(os.TempDir(), "mylist")
		if err != nil {
			errLogger.Fatalln(err)
		}

		fstate, err := f.Stat()
		if err != nil {
			errLogger.Fatalln(err)
		}
		defer os.Remove(fstate.Name())

		fname := fstate.Name()
		infoLogger.Println("Start assembling video...")
		os.Remove(path.Join(*p, *name+".ts"))
		os.Remove(path.Join(*p, *name+".mp4"))
		absPath, err := getAbsPath(*p)
		if err != nil {
			errLogger.Fatalln(err)
		}
		for i := 0; i <= chunksNb; i++ {
			f.WriteString(fmt.Sprintf("file %s/%s%d.ts\n", absPath, *name, i))
		}
		f.Close()

		args := []string{"-safe", "0", "-f", "concat", "-i", path.Join(os.TempDir(), fname), "-c", "copy", path.Join(*p, *name+".ts")}
		cmd := exec.Command(ffmpegPath, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			errLogger.Fatal(err)
		}

		infoLogger.Println("Converting to mp4...")
		args = []string{"-i", path.Join(*p, *name+".ts"), "-acodec", "copy", "-vcodec", "copy", path.Join(*p, *name+".mp4")}
		cmd = exec.Command(ffmpegPath, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			errLogger.Fatalln(err)
		}

		infoLogger.Printf("Assembling %s to mp4 is done...", *name)
		infoLogger.Println("Cleaning chunks...")
		for i := 0; i <= chunksNb; i++ {
			err := os.Remove(path.Join(*p, *name+fmt.Sprintf("%d.ts", i)))
			if err != nil {
				errLogger.Println(err)
			}
		}
		os.Remove(path.Join(*p, *name+".ts"))
		infoLogger.Println("Done :) ")
	}
}

func run(log *log.Logger, cfg Config, nbChunksChan chan int, errChan chan error, doneChan chan bool) {
	var nbChunks int
	var c *check.CloudfrontChecker
	if cfg.Link == "" || cfg.Name == "" {
		errChan <- ErrMissingArgs
		doneChan <- true
		return
	}
	if cfg.Max <= 0 {
		c = check.NewCloudfrontCheckerWithLog(cfg.Link, math.MaxInt-1, log)
	} else {
		c = check.NewCloudfrontCheckerWithLog(cfg.Link, cfg.Max, log)
	}
	nbChunks, err := c.Check()
	if err != nil {
		errChan <- err
		doneChan <- true
		nbChunksChan <- -1
		return
	}
	nbChunksChan <- nbChunks

	log.Printf("Nb of chunks is %d , from %d to %d", nbChunks+1, 0, nbChunks)
	if cfg.Dwn {
		baseLink := check.GetBaseLink(cfg.Link)
		bd := dwn.NewBulkDownloaderWithLog(log, cfg.Name, ".ts", cfg.Path, errChan, doneChan)
		for i := 0; i <= nbChunks; i++ {
			bd.AddUrl(baseLink + strconv.Itoa(i) + ".ts")
		}
		go bd.Download()
	} else {
		doneChan <- true
	}
}
