package utils

import (
	"encoding/base32"
)

// Base32Decode
func Base32Decode(value string) (string, error) {
	decoded, err := base32.StdEncoding.DecodeString(value)
	if nil != err {
		return "", err
	}

	return string(decoded), nil
}

// Base32NopaddingDecode
func Base32NopaddingDecode(value string) (string, error) {
	var b32NoPadding = base32.StdEncoding.WithPadding(base32.NoPadding)
	decoded, err := b32NoPadding.DecodeString(value)
	if nil != err {
		return "", err
	}

	return string(decoded), nil
}
