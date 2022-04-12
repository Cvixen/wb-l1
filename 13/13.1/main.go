package main

import "fmt"

func main() {
	a := 5
	b := 3
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d\nb = %d\n", a, b)
}
