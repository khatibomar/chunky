package check

import (
	"strings"
)

func GetBaseLink(link string) string {
	return strings.Split(link, "chunked")[0] + "chunked/"
}
