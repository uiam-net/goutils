package utils

import (
	crand "crypto/rand"
	"io"
)

// var idChars = []byte(TxtNumbers + TxtAlphabet)

const (
	TxtNumbers         = "0123456789"
	TxtAlphabetDown    = "abcdefghijklmnopqrstuvwxyz"
	TxtAlphabetUp      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	TxtAlphabet        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	TxtSimpleCharaters = "13467ertyiadfhjkxcvbnERTYADFGHJKXCVBN"
)

// randomDigits returns a byte slice of the given length containing
// pseudorandom numbers in range 0-9. The slice can be used as a captcha
// solution.
func randomDigits(length int) []byte {
	return randomBytesMod(length, 10)
}

// randomBytes returns a byte slice of the given length read from CSPRNG.
func randomBytes(length int) (b []byte) {
	b = make([]byte, length)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
	return
}

// randomBytesMod returns a byte slice of the given length, where each byte is
// a random number modulo mod.
func randomBytesMod(length int, mod byte) (b []byte) {
	if length == 0 {
		return nil
	}
	if mod == 0 {
		panic("captcha: bad mod argument for randomBytesMod")
	}
	maxrb := 255 - byte(256%int(mod))
	b = make([]byte, length)
	i := 0
	for {
		r := randomBytes(length + (length / 4))
		for _, c := range r {
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = c % mod
			i++
			if i == length {
				return
			}
		}
	}
}

func GenRdmNumbers(length int) string {
	var chars = []byte(TxtNumbers)
	b := randomBytesMod(length, byte(len(chars)))

	for i, c := range b {
		b[i] = chars[c]
	}

	return string(b)
}

func GenRdmString(length int) string {
	var chars = []byte(TxtNumbers + TxtAlphabetDown + TxtAlphabetUp)
	b := randomBytesMod(length, byte(len(chars)))

	for i, c := range b {
		b[i] = chars[c]
	}

	return string(b)
}

func GenRdmUpstring(length int) string {
	var chars = []byte(TxtAlphabetUp)
	b := randomBytesMod(length, byte(len(chars)))

	for i, c := range b {
		b[i] = chars[c]
	}

	return string(b)
}

func GenRdmDownstring(length int) string {
	var chars = []byte(TxtAlphabetDown)
	b := randomBytesMod(length, byte(len(chars)))

	for i, c := range b {
		b[i] = chars[c]
	}

	return string(b)
}

func GenRdmStringByChars(length int, chars []byte) string {
	b := randomBytesMod(length, byte(len(chars)))

	for i, c := range b {
		b[i] = chars[c]
	}

	return string(b)
}

// ========================= Utils =================== //

// DefaultString DefaultString
func DefaultString(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}
