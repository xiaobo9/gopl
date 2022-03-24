package utils

import (
	"os"
	"strings"
)

var tempDir = "temp"

func MkTempDir() {
	os.MkdirAll(tempDir, 0700)
}

// create tmep file in dir temp
func CreateTempFile(fileName string) (*os.File, error) {
	MkTempDir()
	fileName = tempDir + "/" + fileName
	file, err := os.Create(fileName)
	return file, err
}

// fix prefix "http://" to url
func FixUrlPrefix(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}
