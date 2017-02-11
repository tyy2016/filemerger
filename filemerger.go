package main

import (
	"filemerger"
	"fmt"
	"os"
)

func Usage() {
	fmt.Println("Usage: filemerger [DIRECTORY]\n" +
		//TODO:"       filemerger [ --help | -v | --version ]\n\n" +
		"A very useful tool filemerger\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		Usage()
	}

	fm := filemerger.NewFileMerger(os.Args[1])
	fm.Run()
}
