package goembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed NGGYU.png
var NGGYU []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("NGGYU_new.png", NGGYU, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
