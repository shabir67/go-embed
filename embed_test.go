package goembed

import (
	_ "embed"
	"fmt"
	"testing"
)

var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}
