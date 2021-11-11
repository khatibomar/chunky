package dwn

type ErrDownload struct {
	uri  string
	file string
	err  error
}

func (e ErrDownload) Error() string {
	return e.err.Error()
}

func (e ErrDownload) URI() string {
	return e.uri
}

func (e ErrDownload) File() string {
	return e.file
}

type downloader interface {
	Download() error
}

type bulkDownloader interface {
	Download() []error
}
