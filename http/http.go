package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const defaultTimeout = 5 * time.Second

func Get(url string) ([]byte, error) {
	return AdvancedGet(url, nil, defaultTimeout)
}

func Post(url string, body []byte) ([]byte, error) {
	return AdvancedPost(url, nil, body, defaultTimeout)
}

func PostJSON(url string, data any) ([]byte, error) {
	body, err := jsoniter.Marshal(data)
	if err != nil {
		return nil, err
	}
	return AdvancedPost(url, map[string]string{"Content-Type": "application/json"}, body, defaultTimeout)
}

func AdvancedGet(url string, header map[string]string, timeout time.Duration) ([]byte, error) {
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

func AdvancedPost(url string, header map[string]string, body []byte, timeout time.Duration) ([]byte, error) {
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
