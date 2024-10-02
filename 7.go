package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

// 1. SafeMap реализует конкурентно безопасную карту с использованием sync.Mutex.
type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

// Write добавляет пару ключ-значение в map, защищенную мьютексом.
func (sm *SafeMap) Write(key string, value int) {
	sm.mu.Lock()         // Захватываем мьютекс перед записью
	defer sm.mu.Unlock() // Освобождаем мьютекс после записи
	sm.m[key] = value
}

// Read читает значение из map по ключу, защищенному мьютексом.
func (sm *SafeMap) Read(key string) (int, bool) {
	sm.mu.Lock()         // Захватываем мьютекс перед чтением
	defer sm.mu.Unlock() // Освобождаем мьютекс после чтения
	val, ok := sm.m[key]
	return val, ok
}

// Пример использования SafeMap с горутинами.
func main1() {
	// Инициализируем SafeMap
	sm := SafeMap{m: make(map[string]int)}

	var wg sync.WaitGroup

	// Параллельная запись в SafeMap с использованием 10 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done() // Обозначаем завершение работы горутины
			key := fmt.Sprintf("key%d", n)
			sm.Write(key, n)
			fmt.Printf("SafeMap - Written: %s = %d\n", key, n)
		}(i)
	}

	wg.Wait()

	// Чтение значений из SafeMap
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value, ok := sm.Read(key)
		if ok {
			fmt.Printf("SafeMap - Read: %s = %d\n", key, value)
		} else {
			fmt.Printf("SafeMap - Key not found: %s\n", key)
		}
	}
}

// ////// 2. Также можно реализовать использование sync.Map, что избавляет от необходимости использовать мьютекс вручную.
func main2() {
	var sm sync.Map

	var wg sync.WaitGroup

	// Параллельная запись в sync.Map с использованием 10 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done() // Обозначаем завершение работы горутины
			key := fmt.Sprintf("key%d", n)
			sm.Store(key, n) // Используем Store для записи данных
			fmt.Printf("sync.Map - Written: %s = %d\n", key, n)
		}(i)
	}

	wg.Wait()

	// Чтение значений из sync.Map
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		if value, ok := sm.Load(key); ok {
			fmt.Printf("sync.Map - Read: %s = %d\n", key, value)
		} else {
			fmt.Printf("sync.Map - Key not found: %s\n", key)
		}
	}
}

// ////// 3. SafeMap с использованием RWMutex для улучшенной производительности при частом чтении.
type SafeMapRW struct {
	mu sync.RWMutex
	m  map[string]int
}

// Write добавляет пару ключ-значение в map, защищенную RW-мьютексом.
func (sm *SafeMapRW) Write(key string, value int) {
	sm.mu.Lock()         // Захватываем мьютекс на запись
	defer sm.mu.Unlock() // Освобождаем мьютекс после записи
	sm.m[key] = value
}

// Read читает значение из map по ключу, используя RLock для улучшения производительности при чтении.
func (sm *SafeMapRW) Read(key string) (int, bool) {
	sm.mu.RLock()         // Захватываем мьютекс на чтение
	defer sm.mu.RUnlock() // Освобождаем мьютекс после чтения
	val, ok := sm.m[key]
	return val, ok
}

// Пример использования SafeMapRW с горутинами.
func main3() {
	// Инициализируем SafeMapRW
	sm := SafeMapRW{m: make(map[string]int)}

	var wg sync.WaitGroup

	// Параллельная запись в SafeMapRW с использованием 10 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done() // Обозначаем завершение работы горутины
			key := fmt.Sprintf("key%d", n)
			sm.Write(key, n)
			fmt.Printf("SafeMapRW - Written: %s = %d\n", key, n)
		}(i)
	}

	wg.Wait()

	// Параллельное чтение значений из SafeMapRW
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			value, ok := sm.Read(key)
			if ok {
				fmt.Printf("SafeMapRW - Read: %s = %d\n", key, value)
			} else {
				fmt.Printf("SafeMapRW - Key not found: %s\n", key)
			}
		}(i)
	}

	wg.Wait()
}

// Главная функция для запуска всех вариантов
func main() {
	fmt.Println("Running SafeMap with Mutex")
	main1()
	time.Sleep(2 * time.Second)

	fmt.Println("\nRunning sync.Map")
	main2()
	time.Sleep(2 * time.Second)

	fmt.Println("\nRunning SafeMap with RWMutex")
	main3()
}
