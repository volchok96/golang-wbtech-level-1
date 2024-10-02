package main

import (
	"fmt"
	"log"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// Установка i-го бита в 1
func setBit(n int64, i uint) int64 {
	// Сдвигаем 1 влево на i позиций и выполняем битовую операцию ИЛИ
	return n | (1 << i)
}

// Установка i-го бита в 0
func clearBit(n int64, i uint) int64 {
	// Сдвигаем 1 влево на i позиций и выполняем битовую операцию AND NOT (или AND с отрицанием)
	return n &^ (1 << i)
}

// Проверка i-го бита
func checkBit(n int64, i uint) bool {
	// Сдвигаем 1 влево на i позиций и проверяем результат операции И
	return n&(1<<i) != 0
}

// Основная функция
func main() {
	var num int64 = 0 // Начальное значение - все биты равны 0

	// Индексы битов для установки и очистки
	setIndexes := []uint{3, 5}
	clearIndexes := []uint{3, 5}

	// Проверяем правильность введенных индексов
	for _, i := range setIndexes {
		if i >= 64 {
			log.Fatalf("Ошибка: индекс %d выходит за пределы диапазона 0-63\n", i)
		}
	}

	// Устанавливаем биты по индексам
	for _, i := range setIndexes {
		num = setBit(num, i)
		fmt.Printf("После установки %d-го бита в 1: %064b\n", i, num)
	}

	// Проверяем, установлены ли биты
	for _, i := range setIndexes {
		if checkBit(num, i) {
			fmt.Printf("%d-й бит установлен (1)\n", i)
		} else {
			fmt.Printf("%d-й бит не установлен (0)\n", i)
		}
	}

	// Очищаем биты по индексам
	for _, i := range clearIndexes {
		num = clearBit(num, i)
		fmt.Printf("После установки %d-го бита в 0: %064b\n", i, num)
	}

	// Проверяем, очищены ли биты
	for _, i := range clearIndexes {
		if !checkBit(num, i) {
			fmt.Printf("%d-й бит очищен (0)\n", i)
		} else {
			fmt.Printf("%d-й бит все еще установлен (1)\n", i)
		}
	}
}
