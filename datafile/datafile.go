package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetArr() [10]int {
	arr := [10]int{}
	return arr
}

func GetFloats(filename string) ([]float64, error) {
	var nums []float64

	file, err := os.Open(filename)
	if err != nil {
		//log.Fatal("Cannot read file.", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}

		nums = append(nums, val)

		i++
	}

	if scanner.Err() != nil {
		log.Fatal("Error reading file.", scanner.Err())
	}

	err = file.Close()
	if err != nil {
		log.Fatal("Cannot close file.", err)
	}
	fmt.Println("Nums from file : ", nums)

	return nums, err
}
