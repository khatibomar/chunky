package downloader

type downloader interface {
	Download() error
}
