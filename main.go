package main

import (
	"fmt"

	"chunky.github.com/checker"
)

const (
	u = "https://d2nvs31859zcd8.cloudfront.net/8800ebf66a27286d9a6a_handmade_hero_40088786859_1634093554/chunked/0.ts"
)

func main() {
	length, _ := checker.GetChunksLengthWithMax(u, 10000)
	fmt.Println(length)
}
