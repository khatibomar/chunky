package dwn

import (
	"io"
	"log"
	"strconv"
	"sync"
)

type BulkDownloader struct {
	Urls      []string
	Prefix    string
	Extension string
	Path      string
	Log       *log.Logger
	Routines  int
	ErrChan   chan error
	DoneChan  chan bool
}

func NewBulkDownloader(prefix, extension, path string, routines int, errChan chan error, doneChan chan bool) *BulkDownloader {
	return &BulkDownloader{
		Urls:      nil,
		Prefix:    prefix,
		Extension: extension,
		Path:      path,
		Routines:  routines,
		Log:       log.New(io.Discard, "", 0),
		ErrChan:   errChan,
		DoneChan:  doneChan,
	}
}

func NewBulkDownloaderWithLog(log *log.Logger, prefix, extension, path string, routines int, errChan chan error, doneChan chan bool) *BulkDownloader {
	b := NewBulkDownloader(prefix, extension, path, routines, errChan, doneChan)
	b.Log = log
	return b
}

func (bd *BulkDownloader) AddUrl(url string) {
	bd.Urls = append(bd.Urls, url)
}

func (bd *BulkDownloader) Download() error {
	var fd downloader
	var wg sync.WaitGroup

	jobs := make(chan downloader, len(bd.Urls))
	wg.Add(len(bd.Urls))

	for w := 1; w <= bd.Routines; w++ {
		go worker(&wg, jobs, bd.ErrChan)
	}

	for i, url := range bd.Urls {
		name := bd.Prefix + strconv.Itoa(i) + bd.Extension
		fd = NewFileDownloaderWithLog(bd.Log, url, name, bd.Path)
		jobs <- fd
	}

	wg.Wait()

	close(jobs)

	bd.DoneChan <- true
	return nil
}

func worker(wg *sync.WaitGroup, jobs <-chan downloader, errChan chan<- error) {
	for j := range jobs {
		err := j.Download()
		if err != nil {
			errChan <- err
		}
		wg.Done()
	}
}
