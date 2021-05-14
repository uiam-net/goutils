package hash

import (
	"testing"

	applog "github.com/uiam-net/goutils/logger"
)

func TestGenPsw(t *testing.T) {
	code1, _ := GenPasswordHash("59a185e0865311eb87ea2cfda1b9f6e5-8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92")

	applog.Error(code1)
}
