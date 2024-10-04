package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). 
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func UniqCharacters(str string) bool {
	lowStr := strings.ToLower(str) // Приведение строки к нижнему регистру
	mRunes := make(map[rune]struct{}, len(str)) // Создание карты для проверки уникальности

	for _, char := range lowStr {
		if _, exists := mRunes[char]; exists {
			return false 
		}
		mRunes[char] = struct{}{} 
	}
	return true 
}

func main() {
	tests := []string{
		"abcd",       // true
		"abCdefAaf",  // false
		"aabcd",      // false
		"ABCDEF",     // true
		"abCDab",     // false
	}

	for _, str := range tests {
		fmt.Printf("Уникальные ли буквы в строке \"%s\": %t\n", str, UniqCharacters(str))
	}
}
