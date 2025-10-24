package search

import (
	"unicode"
	"usamaqaisrani/gorep/src/core/pattern"
)

func Search(p pattern.Pattern, line string) bool {	
	if p == pattern.SingleDigit {
		for _, char	 := range line {
			if unicode.IsDigit(char) {
				return true
			}
		}
		return false
	}
	return false
}
