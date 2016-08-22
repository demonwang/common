package test

import (
	"testing"
	"github.com/demonwang/common"
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
	var src *string
	if common.IsNil(src) {
		t.Log("src is nil")
	}else{
		t.Fatal("src is not nil")
	}
}