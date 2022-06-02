package utils

import (
	"io"
	"net/url"
	"strings"
)

func SafeClose(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

func GetLastURLPath(link string) (string, error) {
	url, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	parts := strings.Split(url.Path, "/")
	seg := parts[len(parts)-1]

	return seg, nil
}
