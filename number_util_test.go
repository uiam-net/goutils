package utils

import (
	"testing"

	applog "github.com/uiam-net/goutils/logger"
)

func TestIsDigitPresent(t *testing.T) {
	code := IsDigitPresent(143, 3)

	applog.Error(code)
}
