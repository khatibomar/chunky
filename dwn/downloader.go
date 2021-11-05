package dwn

type downloader interface {
	Download() error
}
