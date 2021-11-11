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

func main() {
	link := flag.String("url", "", `provide a link that have a chunk , example:
https://d2nvs31859zcd8.cloudfront.net/70c102b5b66dbeac89e4_handmade_hero_40072241627_1633745055/chunked/155.ts
`)
	p := flag.String("dir", "", "specify a download path")
	max := flag.Int("max", -1, "provide the excpected max number of files, zero or negative numbers will be treated as max int")
	dwn := flag.Bool("dwn", true, "by default true , false if you just want to get chunks size without downloading files")

	flag.Parse()

	c := Config{
		Link: *link,
		Max:  *max,
		Path: path.Clean(*p),
		Dwn:  *dwn,
	}

	infoLogger := log.New(os.Stdout, "Info : ", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "Error : ", log.Ldate|log.Ltime)

	errs := run(infoLogger, c)
	if len(errs) > 0 {
		for _, err := range errs {
			if errors.Is(err, ErrMissingArgs) {
				fmt.Println(`
chunky is a tool that will allow you to download subscribers only videos from twitch.
The tool is under current developement so many bugs will occur , and many missing features and many hard coded stuff.
So please if you found any bug or missing features feel free to open an issue on project page:
https://github.com/khatibomar/chunky
			`)
				fmt.Println("Usage: ")
				flag.PrintDefaults()
				return
			}
			errLogger.Printf("main: %s\n", err)
		}
	}
}

func run(log *log.Logger, cfg Config) []error {
	var nbChunks int
	var errs []error

	var c *check.CloudfrontChecker

	if cfg.Link == "" {
		errs = append(errs, ErrMissingArgs)
		return errs
	}
	if cfg.Max <= 0 {
		c = check.NewCloudfrontCheckerWithLog(cfg.Link, math.MaxInt-1, log)
	} else {
		c = check.NewCloudfrontCheckerWithLog(cfg.Link, cfg.Max, log)
	}
	nbChunks, err := c.Check()
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	log.Printf("Nb of chunks from %d to %d", 0, nbChunks+1)
	if cfg.Dwn {
		baseLink := check.GetBaseLink(cfg.Link)
		// TODO(khatibomar): make error logging live instead of waiting the function to return
		bd := dwn.NewBulkDownloaderWithLog(log, "test", ".ts", cfg.Path)
		for i := 0; i <= nbChunks; i++ {
			bd.AddUrl(baseLink + strconv.Itoa(i) + ".ts")
		}
		errs = bd.Download()
	}
	return errs
}
