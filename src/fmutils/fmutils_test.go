package fmutils

import "testing"

//func TestGetFilenameList(t *testing.T) {
//	GetFilenameList("../../file")
//}

//func TestDeleteFile(t *testing.T) {
//	DeleteFile("/Users/baidu/GoglandProjects/filemerger/file/0.file")
//}

func TestMakeSoftLink(t *testing.T) {
	MakeSoftLink("/Users/baidu/GoglandProjects/filemerger/file/ln.file","/Users/baidu/GoglandProjects/filemerger/file/1.file")
}
