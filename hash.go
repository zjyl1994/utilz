package utilz

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
	"hash/fnv"
	"io"
	"os"

	"crypto/sha1"

	"golang.org/x/crypto/bcrypt"
)

// MD5
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

// SHA1
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

// SHA256
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

// Password
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Fnv1
func Fnv1(data []byte) uint64 {
	h := fnv.New64()
	h.Write(data)
	return h.Sum64()
}

func Fnv1String(s string) uint64 {
	return Fnv1([]byte(s))
}

// CRC32
func CRC32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func CRC32String(s string) uint32 {
	return CRC32([]byte(s))
}

func CRC32File(path string) (uint32, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	h := crc32.NewIEEE()
	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}
