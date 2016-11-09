package common

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/skip2/go-qrcode"
	"github.com/tealeg/xlsx"
	"reflect"
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

func DumpModelToxls(sheetName string,v interface{}) (*xlsx.File,error) {

	f := xlsx.NewFile()
	sheet,err := f.AddSheet(sheetName)
	if err != nil {
		return f,err
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "序号"
	s := reflect.TypeOf(v).Elem()
	fmt.Println(s)
	//通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		cell := row.AddCell()
		cell.Value = s.Field(i).Tag.Get("col")
     	}

	value := reflect.ValueOf(v)
	for i := 0;i< value.Len() ; i++ {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = fmt.Sprint(i+1)
		in := value.Index(i)
		for j:= 0;j<in.NumField() ;j++  {
			cell = row.AddCell()
			cell.Value = fmt.Sprint(in.Field(j))
		}
	}
	return f,nil
}

func getStringValue()  {

}