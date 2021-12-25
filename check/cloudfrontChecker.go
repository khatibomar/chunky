package check

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	c := NewCloudfrontChecker(url, max)
	c.Log = log
	return c
}

func NewTestCloudfrontChecker(url string, max int) *CloudfrontChecker {
	c := NewCloudfrontChecker(url, max)
	c.GetStatus = mockGetStatusCode
	return c
}

// GetChunksLength will return the chunks length of the video
// it will use max parameter as highest possible chunk length
func (c *CloudfrontChecker) Check() (int, error) {
	c.Url = GetBaseLink(c.Url) + "index-dvr.m3u8"
	c.Log.Println(c.Url)

	resp, err := http.Get(c.Url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	lines := strings.Split(string(r), "\n")

	if len(lines) < 2 {
		return -1, fmt.Errorf("Excpected 2 or more lines got %d\n", len(lines))
	}
	chunks := strings.TrimSuffix(lines[len(lines)-3], ".ts")
	return strconv.Atoi(chunks)
}
