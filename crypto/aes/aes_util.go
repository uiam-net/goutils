package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	cryptoutil "github.com/uiam-net/goutils/crypto"
)

// ==== Public Operations ===== //
// ==== Public Operations ===== //

// Encrypt Encrypt
func Encrypt(data []byte, key, iv []byte) (string, error) {
	if data == nil {
		return "", errors.New("Encrypt Params error")
	}
	if key == nil {
		return "", errors.New("Encrypt Params error")
	}
	if iv == nil {
		return "", errors.New("Encrypt Params error")
	}

	key = expendKey(key)
	iv = expendKey(iv)

	ckey, err := aes.NewCipher(key)
	if nil != err {
		return "", err
	}

	encrypter := cipher.NewCBCEncrypter(ckey, iv)

	// PKCS7补码
	str := cryptoutil.PKCS7Padding([]byte(data))
	out := make([]byte, len(str))

	encrypter.CryptBlocks(out, str)

	return base64.StdEncoding.EncodeToString(out), nil
}

// Decrypt Decrypt
func Decrypt(base64Str string, key, iv []byte) ([]byte, error) {
	key = expendKey(key)
	iv = expendKey(iv)

	ckey, err := aes.NewCipher(key)
	if nil != err {
		return nil, err
	}

	decrypter := cipher.NewCBCDecrypter(ckey, iv)

	base64In, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(base64In))
	decrypter.CryptBlocks(out, base64In)

	// 去除 PKCS7 补码
	out = cryptoutil.PKCS7Unpadding(out)
	if out == nil {
		return nil, errors.New("PKCS7Unpadding error")
	}

	return out, nil
}

// ==== Private Operations ===== //
// ==== Private Operations ===== //

func expendKey(key []byte) []byte {
	for len(key) < 16 {
		key = append(key, key...)
	}

	return key[:16]
}
