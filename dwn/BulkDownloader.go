package dwn

import (
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
}

func NewBulkDownloader(prefix, extension, path string) *BulkDownloader {
	return &BulkDownloader{
		Urls:      nil,
		Prefix:    prefix,
		Extension: extension,
		Path:      path,
		Log:       log.New(io.Discard, "", 0),
	}
}

func NewBulkDownloaderWithLog(log *log.Logger, prefix, extension, path string) *BulkDownloader {
	return &BulkDownloader{
		Urls:      nil,
		Prefix:    prefix,
		Extension: extension,
		Path:      path,
		Log:       log,
	}
}

func (bd *BulkDownloader) AddUrl(url string) {
	bd.Urls = append(bd.Urls, url)
}

func (bd *BulkDownloader) Download() []error {
	var fd downloader
	var errs []error

	for i, url := range bd.Urls {
		fd = NewFileDownloaderWithLog(bd.Log, url, bd.Prefix+strconv.Itoa(i)+bd.Extension, bd.Path)
		if err := fd.Download(); err != nil {
			errs = append(errs, err)
			bd.Log.Println(err)
		}
	}

	return errs
}
