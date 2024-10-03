package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). 
// Символы могут быть unicode.

func flipLine(str string) string {
	// Преобразуем строку в срез рун для поддержки Unicode
	runes := []rune(str)

	// Инвертируем рунный срез на месте, чтобы не выделять лишнюю память
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем обратно в строку и возвращаем результат
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Println(flipLine(str))
}
