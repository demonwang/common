package common

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/skip2/go-qrcode"
)

const (
	TimePattern1 TimePattern = `2006-01-02 15:04:05`
	TimePattern2 TimePattern = `2006-01-02 15:04`
	TimePattern3 TimePattern = `2006-01-02`
	TimePattern4 TimePattern = `2006/01/02`
	TimePattern5 TimePattern = `2006/01/02 15:04:05`
)

type TimePattern string

func IsBlank(src string) bool {
	if len(src) <= 0 {
		return true
	}
	return false
}

func IsNil(src interface{}) bool {
	if src == nil {
		return true
	}
	return false
}

func MakeCaptcha() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func MakeQrCode(src string) ([]byte,error) {
	return qrcode.Encode(src, qrcode.Medium, 256)
}

func FormatTime(time *time.Time,pattern TimePattern) string {
	return time.Format(pattern)
}