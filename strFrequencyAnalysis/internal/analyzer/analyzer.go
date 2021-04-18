package analyzer

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type wordCount struct {
	word  string
	count int
}

func clearWordsFromPunct(words []string) (clearWords []string) {
	var tmpStr strings.Builder
	for _, word := range words {

		tmpStr.Reset()

		word = strings.ToLower(word)

		for _, r := range word {
			if unicode.IsLetter(r) {
				tmpStr.WriteRune(r)
			}
		}
		clearWords = append(clearWords, tmpStr.String())
	}
	return
}

func cutStringOnWords(str string) []string {
	str = strings.ToLower(str)

	words := strings.Fields(str)

	return clearWordsFromPunct(words)
}

func wordFormCheck(str string) {

}

func Analyzer(str string) {
	words := cutStringOnWords(str)
	var sliceCountWords []wordCount

	sort.Slice(words, func(i, j int) bool {
		return words[i] <= words[j]
	})
	for _, wordColl := range words {
		idx := sort.Search(len(sliceCountWords), func(i int) bool {
			return sliceCountWords[i].word == wordColl
		})
		// if not found
		if idx == len(sliceCountWords) {
			sliceCountWords = append(sliceCountWords, wordCount{word: wordColl, count: 1})
		} else {
			sliceCountWords[idx].count += 1
		}
	}

	for _, counted := range sliceCountWords {
		fmt.Printf("%s = %d\n", counted.word, counted.count)
	}
}
