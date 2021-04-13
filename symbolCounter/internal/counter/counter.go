package counter

import (
	"errors"
	"unicode"
)

func checkForLetters(str *string) error {
	//check if string has letters
	for _, runeVal := range *str {
		if unicode.IsLetter(runeVal) {
			return nil
		}
	}
	return errors.New("No letters in string.")
}

func CountLetters(str string) (error, string) {
	return checkForLetters(&str), ""
}
