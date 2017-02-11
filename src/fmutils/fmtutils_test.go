package fmutils

import (
	"testing"
	"fmt"
)

func TestGetFileMd5(t *testing.T) {
	digest := GetFileMd5("./types.go")
	fmt.Printf("%v\n", digest)
}
