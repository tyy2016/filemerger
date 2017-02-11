package filemerger

import (
	mylog "github.com/mlycore/tools/mylog"
	"os"
)

var log =  mylog.NewLogger(os.Stderr, 3)

type FileMerger struct {
	Msg string
}

