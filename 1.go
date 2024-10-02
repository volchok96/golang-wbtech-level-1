package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов). 
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

// Структура Human
type Human struct {
    Name string
    Age  int
}

// Метод для структуры Human
func (h *Human) Speak() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Структура Action, в которую встроена структура Human
type Action struct {
    Human // Встраивание структуры Human в Action
}

// Метод для структуры Action
func (a *Action) Work() {
    fmt.Printf("%s is a Golang Dev!\n", a.Name)
}

func main() {
    // Создаем экземпляр Action
    action := Action{
        Human: Human{
            Name: "Katerina",
            Age:  28,
        },
    }

    // Вызываем метод Human через экземпляр Action
    action.Speak() // Output: Hello, my name is Katerina and I am 28 years old.

    // Вызываем метод Action
    action.Work() // Output: Katerina is a Golang Dev!
}
