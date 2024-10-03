package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// partition разделяет массив на две части относительно опорного элемента (pivot)
// и возвращает индекс опорного элемента после разделения.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // выбираем опорный элемент (pivot) в конце массива.
	i := low - 1       // инициализируем индекс меньшего элемента.

	for j := low; j < high; j++ {
		// если текущий элемент меньше опорного.
		if arr[j] < pivot {
			i++                             // увеличиваем индекс меньшего элемента.
			arr[i], arr[j] = arr[j], arr[i] // меняем местами элементы arr[i] и arr[j].
		}
	}
	// меняем местами опорный элемент с элементом после последнего меньшего элемента.
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1 // возвращаем индекс опорного элемента.
}

// quickSort выполняет рекурсивную быструю сортировку.
func quickSort(arr []int, low, high int) {
	if low < high {
		// Разделяем массив на две части и получаем индекс опорного элемента.
		p := partition(arr, low, high)

		// Рекурсивно сортируем левую часть массива.
		quickSort(arr, low, p-1)
		// Рекурсивно сортируем правую часть массива.
		quickSort(arr, p+1, high)
	}
}

// quickSortWrapper — вспомогательная функция для упрощенного вызова быстрой сортировки.
func quickSortWrapper(arr []int) {
	if len(arr) <= 1 {
		return // массивы с 0 или 1 элементом уже отсортированы
	}
	quickSort(arr, 0, len(arr)-1)
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

func main() {
	// Инициализируем массив для быстрой сортировки.
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Before sorting (QuickSort):", arr)

	// Сортируем массив при помощи быстрой сортировки.
	quickSortWrapper(arr)
	fmt.Println("After sorting (QuickSort):", arr)

	// Пример сортировки случайного массива с использованием быстрой сортировки.
	randomArr := generateRandomArray(10)
	fmt.Println("\nRandom array before sorting (QuickSort):", randomArr)
	quickSortWrapper(randomArr)
	fmt.Println("Random array after sorting (QuickSort):", randomArr)

	// Сортировка с использованием встроенной функции sort
	fmt.Println("\nBefore sorting (Built-in sort):", arr)
	sort.Ints(arr)
	fmt.Println("After sorting (Built-in sort):", arr)

	// Сортировка случайного массива с использованием встроенной функции sort
	fmt.Println("\nRandom array before sorting (Built-in sort):", randomArr)
	sort.Ints(randomArr)
	fmt.Println("Random array after sorting (Built-in sort):", randomArr)
}
