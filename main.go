package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path"

	"chunky.github.com/downloader"

	"chunky.github.com/checker"
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

	err := run(infoLogger, c)
	if err != nil {
		if errors.Is(err, ErrMissingArgs) {
			fmt.Println("Usage: ")
			flag.PrintDefaults()
			return
		}
		errLogger.Fatalf("main: %s\n", err)
	}
}

func run(log *log.Logger, cfg Config) error {
	var nbChunks int
	var err error

	var c *checker.CloudfrontChecker

	if cfg.Link == "" {
		return ErrMissingArgs
	}
	if cfg.Max <= 0 {
		c = checker.NewCloudfrontCheckerWithLog(cfg.Link, math.MaxInt-1, log)
	} else {
		c = checker.NewCloudfrontCheckerWithLog(cfg.Link, cfg.Max, log)
	}
	nbChunks, err = c.GetChunksLength()
	if err != nil {
		return err
	}

	log.Printf("Nb of chunks is %d", nbChunks)
	if cfg.Dwn {
		d := downloader.NewFileDownloaderWithLog(log, cfg.Link, "test.ts", cfg.Path)
		err = d.Download()
		if err != nil {
			return err
		}
	}
	return nil
}
