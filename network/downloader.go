package network

import (
	"io"
	"net/http"
)

// DownloadFile returns the requested file as an io.Reader with corresponding close function.
func DownloadFile(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
