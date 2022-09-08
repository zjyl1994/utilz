package utilz

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/subtle"
	"errors"
	"fmt"

	"golang.org/x/crypto/scrypt"
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

const aesNonceSize = 12

func AesEncode(key, plainText []byte) ([]byte, error) {
	nonce := RandBytes(aesNonceSize)

	dk, err := scrypt.Key(key, nonce, 32768, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(dk)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Seal(nonce, nonce, plainText, nil), nil
}

func AesDecode(key, cipherText []byte) ([]byte, error) {
	if len(cipherText) < aesNonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := cipherText[:aesNonceSize], cipherText[aesNonceSize:]

	dk, err := scrypt.Key(key, nonce, 32768, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	c, err := aes.NewCipher(dk)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, nonce, ciphertext, nil)
}
