package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

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

	flag.Parse()

	length, err := run(*link, *max)
	if err != nil {
		if errors.Is(err, ErrMissingArgs) {
			fmt.Println("Usage: ")
			flag.PrintDefaults()
			return
		}
		log.Fatalf("main: %s", err)
	}
	fmt.Printf("the stream have %d chunks\n", length)
}

func run(link string, max int) (int, error) {
	if link == "" {
		return -1, ErrMissingArgs
	}
	if max <= 0 {
		return checker.GetChunksLength(link)
	} else {
		return checker.GetChunksLengthWithMax(link, max)
	}
}
