package checker

import (
	"net/url"
	"path/filepath"
)

func GetBaseLink(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(u.Path)
	u.Path = dir

	return u.String() + "/"
}
