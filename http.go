package utilz

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const defaultHttpTimeout = 5 * time.Second

func HttpGet(url string) ([]byte, error) {
	return HttpAdvancedGet(url, nil, defaultHttpTimeout)
}

func HttpPost(url string, body []byte) ([]byte, error) {
	return HttpAdvancedPost(url, nil, body, defaultHttpTimeout)
}

func HttpPostJSON(url string, data any) ([]byte, error) {
	body, err := jsoniter.Marshal(data)
	if err != nil {
		return nil, err
	}
	return HttpAdvancedPost(url, map[string]string{"Content-Type": "application/json"}, body, defaultHttpTimeout)
}

func HttpAdvancedGet(url string, header map[string]string, timeout time.Duration) ([]byte, error) {
	hc := http.Client{Timeout: timeout}
	var resp *http.Response
	var err error
	if header == nil {
		resp, err = hc.Get(url)
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		for k, v := range header {
			req.Header.Add(k, v)
		}
		resp, err = hc.Do(req)
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpAdvancedPost(url string, header map[string]string, body []byte, timeout time.Duration) ([]byte, error) {
	hc := http.Client{Timeout: timeout}
	var resp *http.Response
	var err error
	contentType := http.DetectContentType(body)
	if header == nil {
		resp, err = hc.Post(url, contentType, bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
	} else {
		req, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			return nil, err
		}
		for k, v := range header {
			req.Header.Add(k, v)
		}
		req.Header.Add("Content-Type", contentType)
		resp, err = hc.Do(req)
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func Download(url string, timeout time.Duration) ([]byte, error) {
	hc := http.Client{Timeout: timeout}
	resp, err := hc.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func DownloadFile(url, path string, timeout time.Duration) error {
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
