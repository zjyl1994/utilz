package utilz

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
)

func SignECDSA(r io.Reader, privKey []byte) (string, error) {
	// 加载x509格式的私钥文件
	block, _ := pem.Decode(privKey)
	if block == nil {
		return "", errors.New("privKey no pem data found")
	}
	pk, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 对输入进行哈希获取信息摘要
	h := sha256.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}
	hash := h.Sum(nil)
	// ECDSA 签名
	sign, err := ecdsa.SignASN1(rand.Reader, pk, hash)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func VerifyECDSA(r io.Reader, pubKey []byte, sign string) (bool, error) {
	// 加载 x509 格式的公钥
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return false, errors.New("pubKey no pem data found")
	}
	genericPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pk := genericPublicKey.(*ecdsa.PublicKey)
	// 对输入进行哈希获取信息摘要
	h := sha256.New()
	_, err = io.Copy(h, r)
	if err != nil {
		return false, err
	}
	hash := h.Sum(nil)
	// ECDSA 验证
	bSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}
	return ecdsa.VerifyASN1(pk, hash, bSign), nil
}
