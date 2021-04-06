package custom

import (
	"fmt"
)

func Print() {
	fmt.Println("Debug")
}

func CalcValues(num int, num2 int) (sum int) {
	sum = num + num2
	return
}
