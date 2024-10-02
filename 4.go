package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). 
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

// processWorker запускает заданное количество воркеров, которые читают данные из канала и выводят их в stdout.
// Воркеры завершаются, когда контекст завершён.
func processWorker(ctx context.Context, countWorkers int, data <-chan interface{}) {
	var wg sync.WaitGroup

	for i := 1; i <= countWorkers; i++ {
		wg.Add(1)
		workerID := i // Создаем копию индекса, чтобы избежать гонки данных

		go func(id int) {
			defer wg.Done()
			for {
				select {
				case value, ok := <-data:
					if !ok {
						// Канал закрыт, выходим из горутины
						return
					}
					fmt.Printf("Worker %d: Value: %v\n", id, value)
				case <-ctx.Done():
					// Контекст завершён, выходим из горутины
					return
				}
			}
		}(workerID)
	}

	wg.Wait() // Ждём завершения всех воркеров
}

func main() {
	// Ввод количества воркеров
	var numWorkers int
	fmt.Print("Enter number of workers: ")
	_, err := fmt.Scanf("%d", &numWorkers)
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers. Please enter a positive integer.")
		return
	}

	// Канал для передачи данных
	chanData := make(chan interface{})

	// Создаём контекст с отменой для завершения программы
	ctx, cancel := context.WithCancel(context.Background())

	// Генерация данных
	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				// Завершаем генерацию данных, если контекст завершён
				return
			default:
				// Отправляем данные в канал
				chanData <- i
				i++
				time.Sleep(500 * time.Millisecond) // Добавлен sleep для снижения нагрузки
			}
		}
	}()

	// Обработка сигнала завершения (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		<-sigChan
		fmt.Println("\nReceived termination signal, shutting down...")
		cancel()    // Отмена контекста для завершения всех горутин
		close(chanData) // Закрываем канал данных
	}()

	// Запускаем процесс воркеров
	processWorker(ctx, numWorkers, chanData)

	fmt.Println("Program terminated")
}
