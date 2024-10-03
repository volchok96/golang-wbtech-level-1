package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Реализовать бинарный поиск встроенными методами языка.

// binarySearch выполняет бинарный поиск в отсортированном массиве.
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2 // Избегаем переполнения

		// Если элемент найден
		if arr[mid] == target {
			return mid
		}
		// Если элемент меньше среднего, ищем в левой части
		if arr[mid] > target {
			high = mid - 1
		} else { // Если элемент больше среднего, ищем в правой части
			low = mid + 1
		}
	}
	// Элемент не найден
	return -1
}

// generateRandomArray генерирует случайный массив указанного размера.
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Создаем новый генератор случайных чисел
	for i := 0; i < size; i++ {
		arr[i] = rng.Intn(100) // генерируем случайное число от 0 до 99
	}
	return arr
}

// builtinBinarySearch выполняет бинарный поиск с использованием встроенной функции sort.Search.
func builtinBinarySearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
	// Проверяем, найден ли элемент
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1 // элемент не найден
}

// main функция для демонстрации бинарного поиска.
func main() {
	// Генерация и сортировка случайного массива
	randomArr := generateRandomArray(10)
	sort.Ints(randomArr) // Сортируем массив для бинарного поиска
	fmt.Println("Sorted array:", randomArr)

	// Задаем элемент для поиска
	target := randomArr[3] // Выбираем случайный элемент из массива для поиска
	fmt.Println("Searching for:", target)

	// Выполняем бинарный поиск собственной реализацией
	index := binarySearch(randomArr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d (Custom Binary Search)\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array (Custom Binary Search)\n", target)
	}

	// Выполняем бинарный поиск с использованием встроенной функции sort.Search
	indexBuiltin := builtinBinarySearch(randomArr, target)

	if indexBuiltin != -1 {
		fmt.Printf("Element %d found at index %d (Built-in Binary Search)\n", target, indexBuiltin)
	} else {
		fmt.Printf("Element %d not found in the array (Built-in Binary Search)\n", target)
	}
}
