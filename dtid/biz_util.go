package utils

import (
	"strconv"
	"time"

	goutil "github.com/uiam-net/goutils"
)

// OrderBizIDGen BIZIDGen
// 201907301746-544-0298
func genID(format string, nano int, rdm int) string {
	now := time.Now()
	dataStr := now.Format(format)

	secStr := ""
	if nano > 0 {
		nanosec := now.Nanosecond()
		nanosecStr := strconv.Itoa(nanosec)
		secStr = nanosecStr[0:nano]
	}

	rdmStr := ""
	if rdm > 0 {
		rdmStr = goutil.GenRdmNumbers(rdm)
	}

	str := dataStr + secStr + rdmStr

	return str
}

func GenID(nano int, rdm int) string {
	return genID("200601021504", nano, rdm)
}

func GenIDMMDD(nano int, rdm int) string {
	idstr := genID("0601021504", nano, rdm)
	// newIdstr := idstr[2 : len(idstr)-0]

	return idstr
}

// // OrderShortBizIDGenUint64 OrderShortBizIDGenUint64
// func GenShortIDUint64() uint64 {
// 	str := GenShortID()
// 	num := goutil.ParseStrToUint64(str, 0)

// 	return num
// }
