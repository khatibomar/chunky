package dwn

import (
	"errors"
	"fmt"
)

var (
	ErrFileExist = errors.New("File already exists.")
)

type ErrDownload struct {
	Uri  string
	File string
	Err  error
}

func (e ErrDownload) Error() string {
	return fmt.Sprintf("%s: %s\n", e.File, e.Err.Error())
}

type downloader interface {
	Download() error
}

type bulkDownloader interface {
	Download() []error
}
