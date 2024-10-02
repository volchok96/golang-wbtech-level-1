package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	// Используем срез для хранения результатов
	results := make([]int, len(arr))

	wg.Add(len(arr))
	for i, value := range arr {
		// Копируем значение i и value в локальные переменные для замыкания
		go func(i, value int) {
			defer wg.Done()
			results[i] = value * value
		}(i, value)
	}

	wg.Wait()
	fmt.Println(results)
}
