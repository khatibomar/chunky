package checker

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrNotValid = errors.New("the number of chuncks surpass MaxInt")
)

const (
	maxUint = ^uint(0)
	minUint = 0
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

// GetChunksLength will return the chunks length of the video
// it will use MaxInt as highest possible chunk length
func GetChunksLength(link string) (int, error) {
	return GetChunksLengthWithMax(link, maxInt)
}

// GetChunksLength will return the chunks length of the video
// it will use max parameter as highest possible chunk length
func GetChunksLengthWithMax(link string, max int) (int, error) {
	var low int
	var mid int
	var high int
	var currHigh int

	var sub_link string

	low = minUint
	high = max
	currHigh = high

	sub_link = strings.Split(link, "/chunked/")[0] + "/chunked/"
	link = sub_link + strconv.Itoa(high) + ".ts"

	if status, _ := getStatusCode(link); status == 200 {
		return -1, ErrNotValid
	}

	for {
		high = high / 2
		link = sub_link + strconv.Itoa(high) + ".ts"
		log.Println("Trying: " + link)

		if status, _ := getStatusCode(link); status != 200 {
			currHigh = high
		} else {
			high = currHigh
			break
		}
	}
	log.Println("Highest: " + link)

	for {
		mid = (high + low) / 2
		link = sub_link + strconv.Itoa(mid) + ".ts"
		log.Println("Trying: " + link)
		if status, _ := getStatusCode(link); status == 200 {
			low = mid
		} else {
			high = mid
		}
		if high == low+1 {
			break
		}
	}

	return low, nil
}

// getStatusCode will try and make an HTTP Get request to a giving URL
// and return it's HTTP Status Code
func getStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
