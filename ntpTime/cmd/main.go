package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/beevik/ntp"
)

func Count(str string, rn rune) int {
	str = strings.ToLower(str)
	var cnt int
	for _, i := range str {
		if i == rn {
			cnt++
		}
	}

	return cnt
}

func main() {
	time, err := ntp.Time("0.europe.pool.ntp.org")
	if err != nil {
		_, err = io.WriteString(os.Stderr, err.Error())
		if err != nil {
			log.Fatal("All bad.")
		}

		os.Exit(1)
	}
	fmt.Println(time)
	fmt.Printf("%d", Count("Just simple string for count s", 's'))

	os.Exit(0)
}
