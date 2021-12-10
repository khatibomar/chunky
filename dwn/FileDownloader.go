package dwn

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
	fd := NewFileDownloader(url, name, path)
	fd.Log = log
	return fd
}

func (fd *FileDownloader) Download() error {
	p := path.Join(fd.Path, fd.Name)
	fd.Log.Printf("Downloading %s...\n", p)
	_, err := os.Stat(p)
	if !os.IsNotExist(err) {
		return ErrFileExist
	}
	err = os.MkdirAll(fd.Path, 0755)
	if err != nil {
		return err
	}
	resp, err := http.Get(fd.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return os.WriteFile(p, file, 0664)
}
