package check

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CloudfrontChecker struct {
	Url       string
	Max       int
	Log       *log.Logger
	GetStatus func(string) (int, error)
}

func NewCloudfrontChecker(url string, max int) *CloudfrontChecker {
	return &CloudfrontChecker{
		Url:       url,
		Max:       max,
		Log:       log.New(io.Discard, "", 0),
		GetStatus: getStatusCode,
	}
}

func NewCloudfrontCheckerWithLog(url string, max int, log *log.Logger) *CloudfrontChecker {
	return &CloudfrontChecker{
		Url:       url,
		Max:       max,
		Log:       log,
		GetStatus: getStatusCode,
	}
}

func NewTestCloudfrontChecker(url string, max int) *CloudfrontChecker {
	return &CloudfrontChecker{
		Url:       url,
		Max:       max,
		Log:       log.New(io.Discard, "", 0),
		GetStatus: mockGetStatusCode,
	}
}

// GetChunksLength will return the chunks length of the video
// it will use max parameter as highest possible chunk length
func (c *CloudfrontChecker) Check() (int, error) {
	var low int
	var mid int
	var high int
	var currHigh int

	var link string
	var baseLink string

	high = c.Max
	currHigh = high

	baseLink = GetBaseLink(c.Url)
	link = baseLink + strconv.Itoa(high) + ".ts"
	c.Log.Println("Trying: " + link)
	status, err := c.GetStatus(link)
	if err != nil {
		return -1, err
	}
	if status == http.StatusOK {
		status, err = c.GetStatus(baseLink + strconv.Itoa(high+1) + ".ts")
		if err != nil {
			return -1, err
		}
		if status != http.StatusOK {
			return high, nil
		}
		return -1, fmt.Errorf("checker: %s", ErrOverMaxInt)
	}

	for {
		high = high / 2
		link = baseLink + strconv.Itoa(high) + ".ts"
		c.Log.Println("Trying: " + link)

		status, err := c.GetStatus(link)
		if err != nil {
			return -1, err
		}

		if status != http.StatusOK {
			currHigh = high
		} else {
			high = currHigh
			break
		}

		if low >= high {
			return -1, fmt.Errorf("checker: %s", ErrInvalid)
		}
	}
	c.Log.Println("Highest Guess: " + link)

	for {
		mid = (high + low) / 2
		link = baseLink + strconv.Itoa(mid) + ".ts"

		c.Log.Println("Trying: " + link)

		status, err := c.GetStatus(link)
		if err != nil {
			return -1, err
		}
		if status == http.StatusOK {
			low = mid
		} else {
			high = mid
		}
		if high == low+1 {
			break
		}
		if low >= high {
			return -1, fmt.Errorf("checker: %s", ErrUnxcpected)
		}
	}

	return low, nil
}