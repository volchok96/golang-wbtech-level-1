package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func Sleep(seconds int) {
	start := time.Now()

	// Выполняем пустой цикл, пока не пройдет указанное количество секунд
	for time.Since(start) < time.Duration(seconds)*time.Second {
	}
}

func main() {
	fmt.Println("Начало ожидания...")
	Sleep(3) // Ожидание 3 секунды
	fmt.Println("Ожидание завершено.")
}
