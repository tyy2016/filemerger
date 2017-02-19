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
	log.Infof("FileMerger is handling directory: %s", fm.Directory)

	fileList := fm.getFilenameList()
	fm.calMd5(fileList)
}

func (fm *FileMerger) getFilenameList() []string {
	return fmutils.GetFilenameList(fm.Directory)
}

func (fm *FileMerger) calMd5(fileList []string) {
	log.Infoln(len(fileList))
	for _, f := range fileList {
		md5 := fmutils.GetFileMd5(f)
		fm.Count[md5]++
		//TODO: keep the same soft link source
		if md5 != "" {
			if fm.Count[md5] > 1 {
				prev := fm.Filename[md5]
				fm.merge(f, prev)
			}else {
				fm.Filename[md5] = f
			}
		} else {
			fm.Filename[md5] = f
		}
	}
}

// current -> prev
func (fm *FileMerger) merge(current, prev string) {
	log.Infof("merger: current=%s, prev=%s", current, prev)
	fmutils.DeleteFile(current)
	fmutils.MakeSoftLink(current, prev)
}
