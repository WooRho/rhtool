package rfile

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
)

func MoveFile(oldpath, newpath string) error { //跨卷移动
	from, err := syscall.UTF16PtrFromString(oldpath)
	if err != nil {
		return err
	}

	tempNum := 0
	for {
		tmpPath := newpath + ".xlsx"
		_, err = os.Stat(tmpPath)
		if err == nil {
			tempNum++
			newpath = newpath + "(" + strconv.Itoa(tempNum) + ")"
		} else {
			break
		}
	}
	newpath = newpath + ".xlsx"
	to, err := syscall.UTF16PtrFromString(newpath)
	if err != nil {
		return err
	}
	return syscall.MoveFile(from, to) //windows API
}

func FileNameList(filepath string) []string {
	var list []string
	rd, err := ioutil.ReadDir(filepath) //遍历目录
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]is dir\n", fi.Name())

		} else {
			list = append(list, fi.Name())
			fmt.Println(filepath + fi.Name())
		}
	}
	return list
}

func MoveFileList(oldpath, newpath string) {
	for _, fi := range FileNameList(oldpath) { //移动目录的所有文件
		err := MoveFile(oldpath+fi, newpath+fi)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fi + "--Move To -->" + newpath + "--OK!")
	}
}
