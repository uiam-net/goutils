package aes

import (
	"testing"

	applog "github.com/uiam-net/goutils/logger"
)

// TestEncrypt TestEncrypt
func TestEncrypt(t *testing.T) {
	data := "W$EvF1HCAzj2^8=3n5M0u_fO"
	key := "t7rwZ3KNTz+7exQpvK"
	iv := "unile.co"

	dataByte := []byte(data)
	keyByte := []byte(key)
	ivByte := []byte(iv)
	dataEncrypt, err := Encrypt(dataByte, keyByte, ivByte)

	if nil != err {
		applog.Error(err)
	}

	applog.Error(dataEncrypt)
}

// TestDecrypt TestDecrypt
func TestDecryptq(t *testing.T) {
	data := "TGer5j10ZHn0adiIH8Xjhw=="
	key := "abcdxxx"
	iv := "14265d0a68"

	// dataByte := []byte(data)
	keyByte := []byte(key)
	ivByte := []byte(iv)
	code, err := Decrypt(data, keyByte, ivByte)

	if nil != err {
		applog.Error(err)
	}

	applog.Error(string(code))
}
