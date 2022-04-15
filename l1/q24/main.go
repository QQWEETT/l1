package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) *Point {
	return &Point{x: x,
		y: y}
}

// Создаем функцию для нахождения расстояния между двумя точками по формуле квадратный корень((х1-х2)² + (y1-y2)²)
func res(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))

}

func main() {
	p1 := newPoint(12, 17)
	p2 := newPoint(-10, 31)
	resultat := res(p1, p2)
	fmt.Println(resultat)

}
