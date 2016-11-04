package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
)

const (
	//TimePattern1 2006-01-02 15:04:05
	TimePattern1 string = `2006-01-02 15:04:05`
	//TimePattern2 2006-01-02 15:04
	TimePattern2 string = `2006-01-02 15:04`
	//TimePattern3 2006-01-02
	TimePattern3 string = `2006-01-02`
	//TimePattern4 2006/01/02
	TimePattern4 string = `2006/01/02`
	//TimePattern5 2006/01/02 15:04:05
	TimePattern5 string = `2006/01/02 15:04:05`
)

//IsBlank check is string is blank
func IsBlank(src string) bool {
	if len(src) <= 0 {
		return true
	}
	return false
}

//IsNil check is a interface is nil
func IsNil(src interface{}) bool {
	if src == nil {
		return true
	}
	return false
}

//MakeCaptcha make a captcha
func MakeCaptcha() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

//MakeQrCode make a qrcode
func MakeQrCode(src string) ([]byte, error) {
	return qrcode.Encode(src, qrcode.Medium, 256)
}

//FormatTime Format a time a pattern
func FormatTime(time *time.Time, pattern string) string {
	return time.Format(pattern)
}

func MD5Encode(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
