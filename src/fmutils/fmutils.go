package fmutils

import (
	"fmt"
	"os"
	"crypto/md5"
)

func NullFileMd5 () [md5.Size]byte {
	data := []byte("null file md5")
	return md5.Sum(data)
}

func GetFilenameList(directory string) []string {
	fmt.Println("this")
	log.Infof("this is mushroom, %d", 1)
	log.Errorf("error")

	return nil
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
	log.Infof("GetFileMd5 read %d bytes data, md5=%x", n, checkSum)
	digest := fmt.Sprintf("%x", checkSum)

	return digest
}

func DeleteFile(filename string) {

}

func MakeSoftLink(filename string) {

}
