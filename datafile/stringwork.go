package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitString(split string) (nodeOne string, nodeTwo string) {
	strs := strings.Split(split, ",")

	fmt.Println("first : ", strs[0], "second", strs[1])

}

func ReadString(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error to open file.", fileName)
		return
	}

	buf := bufio.NewScanner(file)

	for buf.Scan() {

	}
}
