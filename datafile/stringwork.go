package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitString(split string) (string, string) {
	strs := strings.Split(split, ",")

	// fmt.Println("first : ", strs[0], "second", strs[1])

	return strs[0], strs[1]
}

func ReadString(fileName string) (map[string]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error to open file.", fileName)
		return nil, err
	}

	buf := bufio.NewScanner(file)

	ipMap := make(map[string]string)

	for buf.Scan() {
		str1, str2 := splitString(buf.Text())
		ipMap[str1] = str2
	}

	for index, param := range ipMap {
		fmt.Println(index, param)
	}

	return ipMap, nil
}
