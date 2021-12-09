package dwn

import (
	"fmt"
	"io"
	"log"
	"strconv"
)

type BulkDownloader struct {
	Urls      []string
	Prefix    string
	Extension string
	Path      string
	Log       *log.Logger
	ErrChan   chan error
	InfoChan  chan string
}

func NewBulkDownloader(prefix, extension, path string, errChan chan error) *BulkDownloader {
	return &BulkDownloader{
		Urls:      nil,
		Prefix:    prefix,
		Extension: extension,
		Path:      path,
		Log:       log.New(io.Discard, "", 0),
		ErrChan:   errChan,
		InfoChan:  make(chan string),
	}
}

func NewBulkDownloaderWithLog(log *log.Logger, prefix, extension, path string, errChan chan error) *BulkDownloader {
	b := NewBulkDownloader(prefix, extension, path, errChan)
	b.Log = log
	return b
}

func (bd *BulkDownloader) AddUrl(url string) {
	bd.Urls = append(bd.Urls, url)
}

func (bd *BulkDownloader) Download() error {
	var fd downloader

	for i, url := range bd.Urls {
		name := bd.Prefix + strconv.Itoa(i) + bd.Extension
		fd = NewFileDownloaderWithLog(bd.Log, url, name, bd.Path)
		if err := fd.Download(); err != nil {
			bd.ErrChan <- fmt.Errorf("%s: %w", name, err)
		}
	}
	// TODO(khatibomar): Is this right thing to do?!!!
	close(bd.ErrChan)
	return nil
}
