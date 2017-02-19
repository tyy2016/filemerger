package fmutils

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

)

func NullFileMd5() [md5.Size]byte {
	data := []byte("null file md5")
	return md5.Sum(data)
}

func GetFilenameList(directory string) []string {
	err := os.Chdir(directory)
	if err != nil{
		log.Errorf("Open dir failed! Err:%s",err)
	}else{
		log.Infof("Open %s success",directory)
	}

	dirName, err1 := ioutil.ReadDir(directory)
	if err1 != nil {
		log.Errorf("ReadDir %s failed, Error=%s", directory, err1)
		os.Chdir("..")
		return nil
	}
	log.Infof("CurrentDir is %s", directory)

	filenamelist := make([]string, 0)
	for i := 0; i < len(dirName); i++ {
		pwd, _ := os.Getwd()
		absFilePath := pwd + "/" + dirName[i].Name()
		log.Infof("file:%s", absFilePath)

		if dirName[i].IsDir() {
			if dirName[i].Name() == "." || dirName[i].Name() == ".."{

			} else {
				subFileList := GetFilenameList(absFilePath)
				filenamelist = append(filenamelist, subFileList...)
				os.Chdir("..")
			}

		} else {
			filenamelist = append(filenamelist, absFilePath)
			//fmt.Println(filenamelist)
		}
	}

	log.Infof("FilenameList: len=%d", len(filenamelist))
	return filenamelist
}

func GetFileMd5(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Errorf("GetFileMd5 Error: err=%s", err)
		return ""
	}

	fstat, err1 := f.Stat()
	if err1 != nil {
		log.Errorf("GetFileMd5 Error: err=%s", err1)
		return ""
	}

	content := make([]byte, int(fstat.Size()))
	n, err2 := f.Read(content)
	if err2 != nil {
		log.Errorf("GetFileMd5 Error: err=%s", err2)
		return ""
	}

	checkSum := md5.Sum(content)
	log.Infof("GetFileMd5 %s read %d bytes data, md5=%x", filename, n, checkSum)
	digest := fmt.Sprintf("%x", checkSum)

	return digest
}

func DeleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Errorf("Remove file failed ! %s", err)
	} else {
		log.Infof("Remove %s success", filename)
	}
}

// soft link: dst -> src
func MakeSoftLink(current, prev string) {
	err := os.Symlink(prev, current)
	if err != nil{
		log.Errorf("MakeSoftLink Error: error=%s", err)
	}else{
		log.Infof("Creat link success")
	}
}

