package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, 
// а с другой стороны канала — читать. 
// По истечению N секунд программа должна завершаться.

func main() {
	var n int

	// Запрашиваем у пользователя количество секунд (положительное число)
	for {
		fmt.Print("Enter number of seconds to run the program (positive integer): ")
		_, err := fmt.Scanf("%d", &n)
		if err != nil || n <= 0 {
			// Если ввод не является положительным целым числом, выводим сообщение об ошибке
			fmt.Println("Invalid input. Please enter a positive integer.")
			return
		} else {
			break
		}
	}
	// Создаем канал для передачи данных
	chanI := make(chan int)

	// Канал done используется для сигнализации о завершении программы
	done := make(chan struct{})

	// Горутина для записи данных в канал
	go func() {
		i := 0
		for {
			select {
			case <-done:
				// Если получаем сигнал завершения, закрываем канал данных и выходим из горутины
				close(chanI)
				return
			default:
				// Отправляем текущее значение i в канал
				chanI <- i
				i++
				// Добавляем задержку в 1 секунду между отправками
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		for value := range chanI {
			// Печатаем полученное значение
			fmt.Println(value)
		}
	}()

	// Устанавливаем таймер на завершение программы через N секунд
	time.AfterFunc(time.Duration(n)*time.Second, func() {
		fmt.Println("Time's up!")
		// Закрываем канал done, чтобы сигнализировать об окончании работы
		close(done)
	})

	// Ожидаем закрытия канала done для завершения работы main
	<-done
	fmt.Println("Program terminated gracefully.")
}
