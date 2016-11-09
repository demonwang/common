package test

import (
	"testing"
	"github.com/demonwang/common"
	"os"
)

func Test_isBlank(t *testing.T) {
	src := ""
	if common.IsBlank(src) {
		t.Log("src is blank")
	}else{
		t.Fatal("src is not blank")
	}
}

func Test_isNil(t *testing.T) {
	var src interface{}
	if common.IsNil(src) {
		t.Log("src is nil")
	}else{
		t.Fatal("src is not nil")
	}
}

func Test_makeCaptcha(t *testing.T)  {
	captcha := common.MakeCaptcha()
	if len(captcha) == 6{
		t.Log("make captcha success:",captcha)
	}else{
		t.Fatal("make captcha fail:",captcha)
	}
}

func Test_makeQrCode(t *testing.T)  {
	src := "http://www.baidu.com"
	data,err := common.MakeQrCode(src)
	if err != nil {
		t.Fatal(err)
	}else{
		t.Log("make qrcode ok and save to file")
		f,_:=os.Create(common.MakeCaptcha()+".png")
		f.Write(data)
		f.Close()
	}
}

type TestSt struct {
	Name string `col:"名字"`
	Age  int `col:"年龄"`
	FFF float64 `col:"杀掉"`
}

func Test_DumpModelToXls(t *testing.T)  {
	a := []TestSt{{"a",16,10.2},{"a",16,10.2},{"a",16,10.2}}

	f,err := common.DumpModelToxls("haha",a)
	if err != nil {
		t.Error(err)
		return
	}
	file := os.NewFile(100,"file.xlsx")
	defer file.Close()
	f.Write(file)
}