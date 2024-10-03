package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. 
// По завершению программа должна выводить итоговое значение счетчика.

// Counter - структура-счетчик
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment - метод для инкрементирования счетчика
func (c *Counter) Increment() {
	c.mu.Lock()         // блокируем мьютекс
	defer c.mu.Unlock() // разблокируем мьютекс при выходе из функции
	c.value++           // увеличиваем значение счетчика
}

// Value - метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()         // блокируем мьютекс
	defer c.mu.Unlock() // разблокируем мьютекс при выходе из функции
	return c.value      // возвращаем текущее значение счетчика
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Количество горутин
	numGoroutines := 100

	// Запуск горутин
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()     
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Итоговое значение счетчика
	fmt.Printf("Final counter value: %d\n", counter.Value())
}