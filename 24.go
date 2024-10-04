package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, 
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

// Инкапсулированные (неэкспортируемые параметры)
type Point struct {
	x, y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Формула евклидова расстояния
func DistanceBetweenPoints(p1, p2 *Point) float64 {
	x := p2.x - p1.x
	y := p2.y - p1.y
	return math.Sqrt(x*x + y*y)
}

// Вычисления непосредственно на объекте Point
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	a := NewPoint(1.5, 2.5)
	b := NewPoint(4.5, 6.5)
	dist1 := DistanceBetweenPoints(a, b)
	dist2 := b.Distance(a)

	fmt.Println("1. Distance between points: ", dist1)
	fmt.Println("2. Distance between points: ", dist2)
}