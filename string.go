package common

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/skip2/go-qrcode"
)

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