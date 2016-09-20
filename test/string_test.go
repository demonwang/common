package test

import (
	"os"
	"testing"
	"time"

	"github.com/demonwang/common"
)

func Test_isBlank(t *testing.T) {
	src := ""
	if common.IsBlank(src) {
		t.Log("src is blank")
	} else {
		t.Fatal("src is not blank")
	}
}

func Test_isNil(t *testing.T) {
	var src interface{}
	if common.IsNil(src) {
		t.Log("src is nil")
	} else {
		t.Fatal("src is not nil")
	}
}

func Test_makeCaptcha(t *testing.T) {
	captcha := common.MakeCaptcha()
	if len(captcha) == 6 {
		t.Log("make captcha success:", captcha)
	} else {
		t.Fatal("make captcha fail:", captcha)
	}
}

func Test_makeQrCode(t *testing.T) {
	src := "http://www.baidu.com"
	data, err := common.MakeQrCode(src)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("make qrcode ok and save to file")
		f, _ := os.Create(common.MakeCaptcha() + ".png")
		f.Write(data)
		f.Close()
	}
}

func Test_timeFormat(t *testing.T) {
	timeNow := time.Now()
	stringDate := common.FormatTime(&timeNow, common.TimePattern1)
	if common.IsBlank(stringDate) {
		t.Fatal(stringDate)
	} else {
		t.Log(stringDate)
	}
}
