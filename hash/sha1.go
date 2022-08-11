package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func SHA1(data []byte) string {
	h := sha1.New()
	h.Write(data)
	result := h.Sum(nil)
	return hex.EncodeToString(result)
}

func SHA1String(s string) string {
	return SHA1([]byte(s))
}

func SHA1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
