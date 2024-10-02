package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// ////////// 1. Использование канала для сигнала о завершении работы:
func worker1(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Worker 1: Goroutine stopped")
			return
		default:
			fmt.Println("Worker 1: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main1() {
	stopCh := make(chan struct{})
	go worker1(stopCh)

	time.Sleep(3 * time.Second)
	close(stopCh)
	fmt.Println("Main 1: Sent stop signal to worker 1")
	time.Sleep(1 * time.Second)
}

// ////////// 2. Использование контекста (context.Context):
func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker 2: Goroutine stopped")
			return
		default:
			fmt.Println("Worker 2: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main2() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("Main 2: Sent cancel signal to worker 2")
	time.Sleep(1 * time.Second)
}

// ////////// 3. Использование тикера:
func worker3(tickCountLimit int) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	tickCount := 0

	for {
		select {
		case <-ticker.C:
			tickCount++
			fmt.Printf("Worker 3: Working... Tick %d\n", tickCount)
			if tickCount >= tickCountLimit {
				fmt.Println("Worker 3: Stopping goroutine after reaching tick limit")
				return
			}
		}
	}
}

func main3() {
	go worker3(5)
	time.Sleep(5 * time.Second)
	fmt.Println("Main 3: Worker 3 should have stopped")
}

// ////////// 4. Использование таймера:
func main4() {
	done := make(chan struct{})

	go func() {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Worker 4: Goroutine finished after 2 seconds")
			close(done)
		}
	}()

	<-done
	fmt.Println("Main 4: Main function exiting")
}

// ////////// 5. Использование переменной-флага Atomic:
var stopFlag int32

func worker5() {
	for {
		if atomic.LoadInt32(&stopFlag) == 1 {
			fmt.Println("Worker 5: Goroutine stopped")
			return
		}
		fmt.Println("Worker 5: Working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func main5() {
	go worker5()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stopFlag, 1)
	fmt.Println("Main 5: Set stop flag for worker 5")
	time.Sleep(1 * time.Second)
}

// Главная функция для запуска всех вариантов
func main() {
	fmt.Println("Running main1()")
	main1()
	fmt.Println("\nRunning main2()")
	main2()
	fmt.Println("\nRunning main3()")
	main3()
	fmt.Println("\nRunning main4()")
	main4()
	fmt.Println("\nRunning main5()")
	main5()
}
