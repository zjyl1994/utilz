package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func SHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	result := h.Sum(nil)
	return hex.EncodeToString(result)
}

func SHA256String(s string) string {
	return SHA256([]byte(s))
}

func SHA256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
