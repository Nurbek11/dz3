package dz3

import (
"unicode"
)

func caseChanger(word string) string {
	newWord := ""
	//iterate the word symbols
	for _, r := range word {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			//concatenate the symbols
			newWord += string(unicode.ToLower(r))
		} else {
			//concatenate the symbols
			newWord += string(unicode.ToUpper(r))
		}
	}
	return newWord
}

