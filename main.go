package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
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
https://d2nvs31859zcd8.cloudfront.net/70c102b5b66dbeac89e4_handmade_hero_40072241627_1633745055/chunked/155.ts
`)
	p := flag.String("dir", "", "specify a download path")
	max := flag.Int("max", -1, "provide the excpected max number of files, zero or negative numbers will be treated as max int")
	down := flag.Bool("dwn", true, "by default true , false if you just want to get chunks size without downloading files")

	flag.Parse()

	if *p != "" {
		if (*p)[0] == '~' {
			home, err := os.UserHomeDir()
			if err != nil {
				panic("failed to get home user , this should not happen")
			}
			*p = home + (*p)[1:]
		}
	}

	c := Config{
		Link: *link,
		Max:  *max,
		Path: path.Clean(*p),
		Dwn:  *down,
	}

	infoLogger := log.New(os.Stdout, "Info : ", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "Error : ", log.Ldate|log.Ltime)
	errChan := make(chan error)

	go run(infoLogger, c, errChan)
	// TODO(khatibomar): Do more research is this OK?!
	// should I make an signal close instead of letting caller close the chan?!
	for {
		err, ok := <-errChan
		if !ok {
			break
		}
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
	}

	_, ok := <-errChan
	if ok {
		close(errChan)
	}
}

func run(log *log.Logger, cfg Config, errChan chan error) {
	var nbChunks int
	var c *check.CloudfrontChecker
	if cfg.Link == "" {
		errChan <- ErrMissingArgs
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
		return
	}

	log.Printf("Nb of chunks is %d , from %d to %d", nbChunks+1, 0, nbChunks)
	if cfg.Dwn {
		baseLink := check.GetBaseLink(cfg.Link)
		bd := dwn.NewBulkDownloaderWithLog(log, "", ".ts", cfg.Path, errChan)
		for i := 0; i <= nbChunks; i++ {
			bd.AddUrl(baseLink + strconv.Itoa(i) + ".ts")
		}
		go bd.Download()
	}
}
