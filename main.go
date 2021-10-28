package main

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	u       = "https://d2nvs31859zcd8.cloudfront.net/8800ebf66a27286d9a6a_handmade_hero_40088786859_1634093554/chunked/0.ts"
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

func main() {
	var uu string
	var chunk int = 9999

	for {
		uu = strings.Split(u, "/chunked/")[0] + "/chunked/" + fmt.Sprint(chunk) + ".ts"
		resp, err := http.Get(uu)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			chunk /= 2
			fmt.Printf("Chunk %d : NO\n", chunk)
			continue
		}
		fmt.Printf("chunck %d : exists\n", chunk)
		chunk *= 2
	}

}
