package lib

import (
	"bytes"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "test"
	}
}

func TestBufStr(t *testing.T) {
	strBuf := bytes.NewBufferString("")
	for i := 0; i < 100000; i++ {
		strBuf.WriteString("test")
	}
}

func TestFmtStr(t *testing.T) {
	a, b := "Hello", "world"
	for i := 0; i < 100000; i++ {
		// fmt.Sprintf("%s%s", a, b)
		strings.Join([]string{a, b}, "")
	}
}

func TestData(t *testing.T) {

	for i := 0; i < 100000; i++ {
		var a struct {
			Name string
			Age  int
		}
		a.Name = "Hello"
		a.Age = 10
	}
}

func TestData2(t *testing.T) {
	for i := 0; i < 100000; i++ {
		var a = map[string]interface{}{}
		a["Name"] = "Hello"
		a["Age"] = 10
	}
}
