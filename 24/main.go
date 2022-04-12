/**
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func newPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func distanceBetweenPointers(point1 *Point, point2 *Point) float64 {
	return math.Sqrt(math.Pow(point1.x-point2.x, 2) + math.Pow(point1.y-point2.y, 2))
}
func main() {
	point1 := newPoint(2, 3)
	point2 := newPoint(4, 5)
	fmt.Printf("%4.2f\n", distanceBetweenPointers(point1, point2))
}
