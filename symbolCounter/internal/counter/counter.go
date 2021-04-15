package counter

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type letterSymb struct {
	symbol rune
	count  int
}

func CheckForLetters(str *string) error {
	//check if string has letters
	for _, runeVal := range *str {
		if unicode.IsLetter(runeVal) {
			return nil
		}
	}

	return errors.New("No letters in string.")
}

func PackLetters(str string) (string, error) {
	err := CheckForLetters(&str)
	if err != nil {
		return "", err
	}

	var letters = make(map[rune]int)

	for _, runeVal := range str {
		// check if rune letter
		if unicode.IsLetter(runeVal) {

			// find key in map
			if _, found := letters[runeVal]; !found {
				letters[runeVal] = 1
			} else {
				letters[runeVal]++
			}
		}
	}

	//sort keys
	keys := make([]rune, 0, len(letters))
	for k := range letters {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	//build final string
	var strOut strings.Builder
	for _, k := range keys {
		strOut.WriteRune(k)

		if letters[k] > 1 {
			strOut.WriteString(strconv.Itoa(letters[k]))
		}
	}
	return strOut.String(), nil
}

func UnpackString(str string) (string, error) {
	err := CheckForLetters(&str)
	if err != nil {
		return "", err
	}

	// pack to slice of runes
	runes := []rune(str)

	var sliceLetters []letterSymb

	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			if v := i + 1; v < len(runes) && unicode.IsNumber(runes[i+1]) {
				sliceLetters = append(sliceLetters, letterSymb{symbol: runes[i], count: int(runes[i+1] - '0')})
				i++
			} else {
				sliceLetters = append(sliceLetters, letterSymb{symbol: runes[i], count: 1})
			}
		} else if runes[i] == 92 {
			if v := i + 2; v < len(runes) && unicode.IsNumber(runes[i+2]) {
				sliceLetters = append(sliceLetters, letterSymb{symbol: runes[i+1], count: int(runes[i+2] - '0')})
				i++
			} else {
				sliceLetters = append(sliceLetters, letterSymb{symbol: runes[i+1], count: 1})
			}
		}
	}

	// build string
	var finStr strings.Builder
	for _, val := range sliceLetters {
		for i := 0; i < val.count; i++ {
			finStr.WriteRune(val.symbol)
		}
	}

	return finStr.String(), nil
}
