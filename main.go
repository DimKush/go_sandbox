package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func countAverage(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum / len(arr)
}

func main() {
	rand.Seed(time.Now().Unix())

	randomized := rand.Intn(100) + 1

	fmt.Println()

	for {

		fmt.Print("Enter a number from 1 to 100: ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Incorrect input %s", err)
			continue
		}

		input = strings.TrimSpace(input)
		fmt.Print(input)

		inputNum, err := strconv.Atoi(input)
		if err != nil || inputNum > 100 || inputNum < 0 {
			log.Fatal("Incorrect number")
		}

		if inputNum == randomized {
			fmt.Println("Congratulations, you guessed a randomized number.", randomized)
		} else {
			fmt.Println("Sorry, you couldn't guess a randomized number : ", randomized)
		}

		fmt.Print("Continue? yes/no? : ")
		reader = bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')

		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		if input == "no" {
			break
		} else if input == "yes" {
			randomized = rand.Intn(100) + 1
		}
	}

	arr := []int{10, 23, 455, 332, 1}
	fmt.Print(countAverage(arr))
}
