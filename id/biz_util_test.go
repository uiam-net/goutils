package utils

import (
	"testing"

	goutil "github.com/uiam-net/goutils"
	applog "github.com/uiam-net/goutils/logger"
)

func TestBIZIDGen(t *testing.T) {
	code := GenID(3, 4)

	var a uint64 = goutil.ParseStrToUint64(code, 0)
	applog.Info(a)
}

func TestIDMMDDGen(t *testing.T) {
	code := GenIDMMDD(2, 0)

	var a int64 = goutil.ParseStrToInt64(code, 0)
	applog.Info(a)

	applog.Info(code)
}

func TestShortIDGen(t *testing.T) {
	code := GenIDMMDD(3, 5)

	applog.Info(code)
}

func TestOriIDID(t *testing.T) {
	// code := OrderOriIDRestore("order-2019080812282674502")

	// applog.Info(code)
}
