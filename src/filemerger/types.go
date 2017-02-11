package filemerger

import (
	mylog "github.com/mlycore/tools/mylog"
	"os"
	"report"
)

var log = mylog.NewLogger(os.Stderr, 3)

type FileMerger struct {
	Directory string            // the directory will be dealt with
	Filename  map[string]string // MD5 - Filename mapping
	Count     map[string]int32  // MD5 - Count mapping

	//	Reports   chan report.Report // send report to Reporter
	Reporter report.Reporter    // Reporter will make some report to stdout or file.
}
