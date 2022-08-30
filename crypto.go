package utilz

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
)

func Pkcs7Pad(buf []byte, size int) ([]byte, error) {
	if size < 1 || size > 255 {
		return nil, fmt.Errorf("Pkcs7Pad: inappropriate block size %d", size)
	}
	i := size - (len(buf) % size)
	return append(buf, bytes.Repeat([]byte{byte(i)}, i)...), nil
}

func Pkcs7Unpad(buf []byte) ([]byte, error) {
	if len(buf) == 0 {
		return nil, errors.New("Pkcs7Unpad: bad padding")
	}

	padLen := buf[len(buf)-1]
	toCheck := 255
	good := 1
	if toCheck > len(buf) {
		toCheck = len(buf)
	}
	for i := 0; i < toCheck; i++ {
		b := buf[len(buf)-1-i]

		outOfRange := subtle.ConstantTimeLessOrEq(int(padLen), i)
		equal := subtle.ConstantTimeByteEq(padLen, b)
		good &= subtle.ConstantTimeSelect(outOfRange, 1, equal)
	}

	good &= subtle.ConstantTimeLessOrEq(1, int(padLen))
	good &= subtle.ConstantTimeLessOrEq(int(padLen), len(buf))

	if good != 1 {
		return nil, errors.New("Pkcs7Unpad: bad padding")
	}

	return buf[:len(buf)-int(padLen)], nil
}
