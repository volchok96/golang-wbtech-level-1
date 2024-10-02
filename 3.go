package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

// Версия 1: Использование WaitGroup и Mutex для управления суммированием квадратов чисел в конкурентной среде
func mainMethod1() {
	// Инициализируем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Массив чисел, для которых нужно найти сумму квадратов
	arr := [...]int{2, 4, 6, 8, 10}

	// Инициализируем Mutex для защиты общей переменной от гонок данных
	var mu sync.Mutex

	// Переменная для хранения суммы квадратов
	var sum int = 0

	// Увеличиваем счётчик WaitGroup на количество чисел в массиве
	wg.Add(len(arr))
	for _, value := range arr {
		// Копируем значение, чтобы избежать проблемы с замыканием
		value := value
		go func() {
			defer wg.Done() // Обозначаем, что горутина завершила выполнение
			// Используем Mutex для защиты доступа к переменной `sum`
			mu.Lock()
			sum += value * value
			mu.Unlock()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим результат
	fmt.Printf("Get result using WaitGroup and Mutex...\n\nThe sum of squares is: %d\n", sum)
}

// Функция для вычисления квадрата числа и отправки результата в канал
func square(n int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы функции
	ch <- n * n     // Отправляем результат в канал
}

// Версия 2: Использование WaitGroup и канала для параллельного вычисления квадратов
func mainMethod2() {
	// Массив чисел, для которых нужно найти сумму квадратов
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	// Создаем буферизированный канал, размер которого равен количеству элементов в массиве
	ch := make(chan int, len(numbers))

	// Запускаем горутину для каждого числа в массиве
	for _, n := range numbers {
		wg.Add(1)             // Увеличиваем счётчик WaitGroup перед запуском каждой горутины
		go square(n, &wg, ch) // Запускаем горутину для вычисления квадрата числа
	}

	// Закрываем канал, когда все горутины завершены
	go func() {
		wg.Wait() // Ожидаем завершения всех горутин
		close(ch) // Закрываем канал, чтобы завершить чтение из него
	}()

	// Суммируем результаты из канала
	sum := 0
	for result := range ch {
		sum += result
	}

	// Выводим результат
	fmt.Printf("Get result using WaitGroup and buffered channel...\n\nThe sum of squares is: %d\n", sum)
}

func main() {
	method := flag.String("method", "", "Choose the method: 1 for WaitGroup+Mutex, 2 for WaitGroup+Channel")
	flag.Parse()

	if *method == "" {
		fmt.Println("Error: You must specify the method to use.")
		fmt.Println("Usage:\ngo run 3.go -method=1\nor\ngo run 3.go -method=2")
		os.Exit(1)
	}

	switch *method {
	case "1":
		mainMethod1() // Запускаем первый метод
	case "2":
		mainMethod2() // Запускаем второй метод
	default:
		fmt.Println("Error: Invalid method. Please choose 1 or 2.")
		fmt.Println("Usage:\ngo run 3.go -method=1\nor\ngo run 3.go -method=2")
		os.Exit(1)
	}
}
