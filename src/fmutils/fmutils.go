package fmutils

import (
	"fmt"
	"io/ioutil"
)

func GetFilenameList(directory string) []string {

	log.Infof("this is mushroom, %d", 1)

	dirName, _ := ioutil.ReadDir(directory)
	log.Infof("CurrentDir is %s",directory)


	filenamelist := make([]string, 0)
	for i := 0; i < len(dirName); i++ {
		if dirName[i].IsDir() {

		} else {
			filename := dirName[i].Name()
			filenamelist = append(filenamelist, filename)
		}
	}
	fmt.Println("Filename_list:", filenamelist)

	return nil
}
