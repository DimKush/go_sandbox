package main

import (
	"fmt"
	"strings"
)

// LeetCode

func middleJustify(words []string, diff int, i int, j int) string {
	spacedNeeded := j - i - 1

	spaces := diff / spacedNeeded
	extraSpaces := diff % spacedNeeded

	var strb strings.Builder
	strb.WriteString(words[i])

	for k := i + 1; k < j; k++ { // O(n)
		var tmp int

		if extraSpaces > 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		extraSpaces--

		spacesToApply := spaces + tmp
		strb.WriteString(strings.Repeat(" ", spacesToApply) + words[k])
	}

	return strb.String()
}

func leftJustify(words []string, diff int, i int, j int) string {
	spacesOnRight := diff - (j - i - 1)
	var strb strings.Builder

	strb.WriteString(words[i])

	for k := i + 1; k < j; k++ {
		strb.WriteString(" " + words[k])
	}

	strb.WriteString(strings.Repeat(" ", spacesOnRight))

	return strb.String()
}

func justifyTheTextWords(words []string, maxWidth int) []string {
	var result []string

	i := 0
	for i < len(words) {
		j := i + 1
		lineLength := len(words[i])
		for j < len(words) && (lineLength+len(words[j])+(j-i-1) < maxWidth) {
			lineLength += len(words[j])
			j++
		}
		diff := maxWidth - lineLength
		numOfWords := j - i
		if numOfWords == 1 || j >= len(words) {
			result = append(result, leftJustify(words, diff, i, j))
		} else {
			result = append(result, middleJustify(words, diff, i, j))
		}
		i = j
	}
	for _, val := range result {
		fmt.Println(val)
	}

	return result
}

func main() {
	maxWidth := 16

	var words = []string{"This", "is", "an", "example", "of", "text", "justification."}
	//text := "This is an example of text justification."
	fmt.Println(justifyTheTextWords(words, maxWidth))
}
