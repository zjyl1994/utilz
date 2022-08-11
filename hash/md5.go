package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	result := h.Sum(nil)
	return hex.EncodeToString(result)
}

func MD5String(s string) string {
	return MD5([]byte(s))
}

func MD5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
