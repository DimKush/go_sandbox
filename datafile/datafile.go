package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloats(filename string) ([3]float64, error) {
	var nums [3]float64

	file, err := os.Open(filename)
	if err != nil {
		//log.Fatal("Cannot read file.", err)
		return nums, err
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if i >= len(nums) {
			log.Println("ATTENTION: array override")
			i = 0
		}

		nums[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nums, err
		}

		i++
	}

	if scanner.Err() != nil {
		log.Fatal("Error reading file.", scanner.Err())
	}

	err = file.Close()
	if err != nil {
		log.Fatal("Cannot close file.", err)
	}

	return nums, err
}
