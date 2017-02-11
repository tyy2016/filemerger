package report

import (
	"github.com/mlycore/tools/mylog"
	"os"
)

var log = mylog.NewLogger(os.Stderr, mylog.INFO)

type Report struct {
	Id       int32  `json:"id"`
	Filename string `json:"filename"`
	Md5      string `json:"md5"`
	Status   string `json:"status"` // Soft link or Origin file.
}

type Reporter struct {
	Reports chan Report

	// io.Writer
}

