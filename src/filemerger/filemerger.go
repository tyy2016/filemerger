package filemerger

import (
	"fmt"
	"fmutils"
)

func NewFileMerger(directory string) *FileMerger {
	/*
	return &FileMerger{
		Directory: directory,
		Filename: make(map[string]string),
		Count: make(map[string]int32),
	}
	*/

	fm := new(FileMerger) // creat a variables
	fm.Directory = directory
	fm.Filename = make(map[string]string)
	fm.Count = make(map[string]int32)

	return fm
}

func (fm *FileMerger) Run() { //fm:recevier *FileMerge:type Run:func
	fmt.Printf("FileMerger is handling directory: %s\n", fm.Directory)
	log.Infof("FileMerger is handling directory: %s\n", fm.Directory)

	fileList := fmutils.GetFilenameList(fm.Directory)

	fm.calMd5(fileList)
}

func (fm *FileMerger) getFilenameList() {
	fmutils.GetFilenameList(fm.Directory)
}

func (fm *FileMerger) calMd5(fileList []string) {
	for _, f := range fileList {
		md5 := fmutils.GetFileMd5(f)
		fm.Count[md5]++
		if md5 != "" {
			fm.Filename[md5] = fm.Directory
			if fm.Count[md5] > 1 {
				fm.merge(f)
			}
		}
	}
}

func (fm *FileMerger) merge(filename string) {
	fmutils.DeleteFile(filename)
	var src, dst string
	fmutils.MakeSoftLink(src, dst)
}
