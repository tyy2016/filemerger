package main

import (
	log "github.com/mlycore/tools/mylog"
	"os"
	"fmt"
	"filemerger"
)

func Usage() {
	fmt.Println("Usage: filemerger [DIRECTORY]\n" +
		       //TODO:"       filemerger [ --help | -v | --version ]\n\n" +
		       "A very useful tool filemerger\n")
	os.Exit(1)
}

func main() {
	var logger = log.NewLogger(os.Stderr, 3)
	logger.Infoln("Hello World")

	if len(os.Args) != 2 {
		Usage()
	}

	fp := new(filemerger.FileMerger)
	fp.Run()
}