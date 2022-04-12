/**
Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func main() {
	a := 5
	b := 3
	a, b = b, a
	fmt.Printf("a = %d\nb = %d\n", a, b)
}
