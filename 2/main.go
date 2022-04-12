/**
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import "fmt"

func sqrt(x int, c chan int) {
	c <- x * x
}

func main() {
	digits := [5]int{2, 4, 6, 8, 10}

	c := make(chan int)
	for i := 0; i < len(digits); i++ {
		go sqrt(digits[i], c)
		fmt.Println(<-c)
	}
}
