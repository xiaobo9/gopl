package ssh

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func readDictFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			result = append(result, passwd)
		}
	}
	return result, err
}

func checkAlive(ip string) bool {
	alive := false
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ip, "22"), time.Second*30)
	if err == nil {
		alive = true
	}
	return alive
}
