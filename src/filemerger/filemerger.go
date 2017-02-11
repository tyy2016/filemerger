package filemerger

import (
	"fmt"
	"fmutils"
)

func (fp *FileMerger) Run() {
	fmt.Println("FileMerger is Running")
	log.Infof("FileMerger is Running")
	fmutils.GetFilenameList("hello")
}

