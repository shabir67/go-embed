package goembed

import (
	"embed"
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

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))

}
