package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		errStr := err.Error()
		os.Stderr.WriteString(errStr)
		os.Stderr.WriteString("\nExit with return status 1\n")
		os.Exit(1)
	}
	fmt.Println(time)

}
