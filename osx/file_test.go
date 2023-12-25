package osx_test

import (
	"github.com/ice-cream-heaven/utils/osx"
	"os"
	"testing"
)

func TestAppend(t *testing.T) {
	path := "E:\\ice-cream-heaven\\utils\\osx\\tmp.txt"
	err := os.WriteFile(path, []byte("123"), 0666)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	err = osx.Append(path, []byte("123"))
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	t.Log(string(buf))
}
