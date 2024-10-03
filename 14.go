package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

// Проверка типа через type switch
func identifyTypeWithSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("[Switch] The type of the variable is int and the value is %d\n", v)
	case string:
		fmt.Printf("[Switch] The type of the variable is string and the value is %s\n", v)
	case bool:
		fmt.Printf("[Switch] The type of the variable is bool and the value is %t\n", v)
	case chan int:
		fmt.Printf("[Switch] The type of the variable is channel of int\n")
	default:
		fmt.Printf("[Switch] Unknown type\n")
	}
}

// Проверка типа через рефлексию
func identifyTypeWithReflect(i interface{}) {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Int:
		fmt.Printf("[Reflect] The type of the variable is int and the value is %d\n", i)
	case reflect.String:
		fmt.Printf("[Reflect] The type of the variable is string and the value is %s\n", i)
	case reflect.Bool:
		fmt.Printf("[Reflect] The type of the variable is bool and the value is %t\n", i)
	case reflect.Chan:
		fmt.Printf("[Reflect] The type of the variable is channel\n")
	default:
		fmt.Printf("[Reflect] Unknown type: %s\n", t.Kind())
	}
}

// Основная функция
func main() {
	// Объявление переменных
	var a int = 42
	var b string = "hello"
	var c bool = true
	var d chan int = make(chan int)
	var e float64 = 1.28 // Пример неизвестного типа

	// Создаем срез переменных для удобного прохода через обе функции
	variables := []interface{}{a, b, c, d, e}

	fmt.Println("=== Identifying types with type switch ===")
	for _, v := range variables {
		identifyTypeWithSwitch(v)
	}

	fmt.Println("\n=== Identifying types with reflection ===")
	for _, v := range variables {
		identifyTypeWithReflect(v)
	}
}
