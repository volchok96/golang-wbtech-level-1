package main

import "fmt"

// Удалить i-ый элемент из слайса.

// Функция для удаления элемента с заданным индексом из слайса
func remove(slice []int, i int) []int {
    // Проверка, что индекс находится в допустимом диапазоне
    if i < 0 || i >= len(slice) {
        fmt.Println("Индекс вне допустимого диапазона")
        return slice
    }

    // Создаем новый слайс, чтобы не изменять исходный слайс
    newSlice := make([]int, len(slice)-1)
    copy(newSlice, slice[:i])
    copy(newSlice[i:], slice[i+1:])
    
    return newSlice
}

func main() {
    // Пример слайса
    slice := []int{1, 2, 3, 4, 5}

    // Индекс элемента для удаления
    i := 2

    // Удаление элемента
    newSlice := remove(slice, i)

    // Вывод результата
    fmt.Println("Original slice:", slice)
    fmt.Println("New slice after removing element at index", i, ":", newSlice)

    // Пример использования с некорректным индексом
    invalidIndex := 10
    newSlice = remove(slice, invalidIndex)

    // Вывод результата
    fmt.Println("Attempt to remove element at invalid index", invalidIndex, ":", newSlice)
}
