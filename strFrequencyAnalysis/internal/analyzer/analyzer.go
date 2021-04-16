package analyzer

import (
	"fmt"
	"strings"
)

func cutStringOnWords(str string) {
	str = strings.ToLower(str)

	words := strings.Fields(str)

	for _, val := range words {
		fmt.Println(val)
	}
}

func Analyzer(str string) {
	cutStringOnWords(str)
}
