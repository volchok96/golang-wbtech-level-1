package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7} // Исходный массив чисел
	first := make(chan int)              // Первый канал для чисел из массива
	second := make(chan int)             // Второй канал для результатов умножения на 2

	var wg sync.WaitGroup

	// Первая горутина: записывает числа из массива в канал `first`
	wg.Add(1)
	go func() {
		defer wg.Done() // Сообщает, что горутина завершила работу
		for _, value := range arr {
			first <- value
		}
		close(first) // Закрываем канал, чтобы сообщить, что больше значений не будет
	}()

	// Вторая горутина: читает значения из `first`, умножает на 2 и записывает в канал `second`
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range first { // Читаем значения из `first`
			second <- value * 2 // Умножаем значение на 2 и отправляем в `second`
		}
		close(second) // Закрываем канал, чтобы сообщить, что больше значений не будет
	}()

	// Третья горутина: читает значения из `second` и выводит их на экран
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range second { // Читаем значения из `second`
			fmt.Println(value) // Выводим результат в стандартный вывод
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}
