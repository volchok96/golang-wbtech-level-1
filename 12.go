package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree),
// создать для нее собственное множество.

func main() {
	// Исходная последовательность строк. Мы имеем повторы: "cat" встречается трижды.
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустую карту, которая будет играть роль множества.
	// Ключи будут строками, а значениями - пустые структуры (struct{}).
	// Использование struct{} позволяет экономить память, так как он не занимает пространство в памяти.
	set := make(map[string]struct{})

	// Проходим по каждой строке в исходной последовательности.
	for _, item := range sequence {
		// Добавляем элемент в карту, где ключ - это строка, а значение - пустая структура.
		// Если строка уже есть в карте, то она не будет добавлена повторно.
		set[item] = struct{}{}
	}

	// Выводим результат.
	fmt.Println("Set:", set)

	// Выводим элементы множества в виде среза для лучшего восприятия
	var result []string
	for key := range set {
		result = append(result, key)
	}

	fmt.Println("Unique Elements in Set:", result)
}
