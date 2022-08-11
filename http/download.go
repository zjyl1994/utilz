package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func DownloadToMem(url string, timeout time.Duration) ([]byte, error) {
	hc := http.Client{Timeout: timeout}
	resp, err := hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func DownloadToFile(url, path string, timeout time.Duration) error {
	hc := http.Client{Timeout: timeout}
	resp, err := hc.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
