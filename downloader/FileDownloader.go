package downloader

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type FileDownloader struct {
	Url  string
	Name string
	Path string
	Log  *log.Logger
}

func NewFileDownloader(url, name, path string) *FileDownloader {
	return &FileDownloader{
		Url:  url,
		Name: name,
		Path: path,
		Log:  log.New(io.Discard, "", 0),
	}
}

func NewFileDownloaderWithLog(log *log.Logger, url, name, path string) *FileDownloader {
	return &FileDownloader{
		Url:  url,
		Name: name,
		Path: path,
		Log:  log,
	}
}

func (fd *FileDownloader) Download() error {
	err := os.MkdirAll(fd.Path, 0755)
	if err != nil {
		return err
	}
	resp, err := http.Get(fd.Url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	fd.Log.Printf("Downloading %s to %s\n", fd.Name, fd.Path)

	file, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	p := path.Join(fd.Path, fd.Name)

	return os.WriteFile(p, file, 0664)
}
