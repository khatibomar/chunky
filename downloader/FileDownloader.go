package downloader

import (
	"io/ioutil"
	"net/http"
	"os"
)

type FileDownloader struct {
	Url  string
	Name string
}

func NewFileDownloader(url, name string) *FileDownloader {
	return &FileDownloader{
		Url:  url,
		Name: name,
	}
}

func (fd *FileDownloader) Download() error {
	resp, err := http.Get(fd.Url)
	defer resp.Body.Close()

	file, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	err = os.WriteFile(fd.Name, file, 0664)
	if err != nil {
		return err
	}
	return nil
}
