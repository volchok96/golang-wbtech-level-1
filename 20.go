package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func reverseWords(str string) string {
	// Разделяем строку на слова
	words := strings.Fields(str)

	// Инвертируем массив слов на месте
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Соединяем инвертированный массив слов в строку
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str)) // Output: sun dog snow
}
