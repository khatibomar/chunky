package main

import (
	"fmt"
	"log"

	"chunky.github.com/checker"
)

const (
	goodLink = "https://d2nvs31859zcd8.cloudfront.net/8800ebf66a27286d9a6a_handmade_hero_40088786859_1634093554/chunked/0.ts"
	badLink  = "https://d2nvs31859zcd8.cloudfront.net/8800ebf21337286d9a6a_handmade_hero_40088786859_1634093554/chunked/0.ts"
)

func main() {
	length, err := checker.GetChunksLengthWithMax(badLink, 10000)
	if err != nil {
		log.Fatalf("main: %s", err)
	}
	fmt.Printf("the stream have %d chunks\n", length)
}
