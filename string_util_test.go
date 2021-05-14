package utils

import (
	"testing"

	applog "github.com/uiam-net/goutils/logger"
)

func TestRandomNumGen(t *testing.T) {
	code := GenRdmNumbers(10)

	applog.Error(code)
}

func TestRandomNumGe2(t *testing.T) {
	code := GenRdmDownstring(10)

	applog.Error(code)
}

func TestRandomNumGen3(t *testing.T) {
	TxtNumbers := "012356789"

	code := GenRdmStringByChars(100, []byte(TxtNumbers))

	applog.Error(string(code))
}
