package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

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
	max := flag.Int("max", -1, "provide the excpected max number of files, zero or negative numbers will be treated as max int")
	dwn := flag.Bool("dwn", true, "by default true , false if you just want to get chunks size without downloading files")
	flag.Parse()

	infoLogger := log.New(os.Stdout, "Info : ", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stderr, "Error : ", log.Ldate|log.Ltime)

	err := run(infoLogger, *link, *max, *dwn)
	if err != nil {
		if errors.Is(err, ErrMissingArgs) {
			fmt.Println("Usage: ")
			flag.PrintDefaults()
			return
		}
		errLogger.Fatalf("main: %s\n", err)
	}
}

func run(log *log.Logger, link string, max int, dwn bool) error {
	var nbChunks int
	var err error

	var c *checker.CloudfrontChecker

	if link == "" {
		return ErrMissingArgs
	}
	if max <= 0 {
		c = checker.NewCloudfrontCheckerWithLog(link, checker.MaxInt, log)
	} else {
		c = checker.NewCloudfrontCheckerWithLog(link, max, log)
	}
	nbChunks, err = c.GetChunksLength()
	if err != nil {
		return err
	}

	log.Printf("Nb of chunks is %d", nbChunks)
	if dwn {
		d := downloader.NewFileDownloader(link, "test.ts")
		d.Download()
	}
	return nil
}
