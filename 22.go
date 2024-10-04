package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b, значение которых > 2^20.

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// Задание значений a и b, больше 2^20
	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s + %s = %s\n", a.String(), b.String(), sum.String())

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Разность: %s - %s = %s\n", a.String(), b.String(), diff.String())

	// Умножение
	product := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s * %s = %s\n", a.String(), b.String(), product.String())

	// Деление
	quotient := new(big.Int).Div(a, b)
	fmt.Printf("Частное: %s / %s = %s\n", a.String(), b.String(), quotient.String())
}
