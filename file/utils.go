package zconverter

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readFile2ByteBuffer(fileName string) ([]byte, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fileSize := fileInfo.Size()
	buf := make([]byte, fileSize)
	count, err := f.Read(buf)
	if err != nil && err != io.EOF {
		log.Println("", count, err)
		return nil, err
	}
	return buf, nil
}

// list dir then do someting with func doSomething
// func doSomething will receive the file info include dir and fileName
func list(path string, doSomething func(dir string, fileName string)) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Panicf("%v", err)
		return
	}
	if !fileInfo.IsDir() {
		doSomething(".", fileInfo.Name())
		return
	}
	listDir(path, doSomething)
}

func listDir(dirPath string, f func(dir string, file string)) {
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		log.Panicf("%v", err)
		return
	}
	if !fileInfo.IsDir() {
		log.Printf("%v/%v, not dir", dirPath, fileInfo.Name())
		f(dirPath, fileInfo.Name())
		return
	}
	log.Printf("%v, dir", dirPath)
	files, _ := ioutil.ReadDir(dirPath)
	for _, fi := range files {
		if fi.IsDir() {
			listDir(dirPath+"/"+fi.Name(), f)
		} else {
			f(dirPath, fi.Name())
		}
	}
}
