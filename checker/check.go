package checker

import (
	"errors"
	"net/http"
)

type logger interface {
	GetChunksLength() (int, error)
}

var (
	ErrInvalid    = errors.New("The provided link is Invalid")
	ErrOverMaxInt = errors.New("The number of chuncks surpass Max Integer")
	ErrUnxcpected = errors.New("Unxcpected error happened")
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

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
