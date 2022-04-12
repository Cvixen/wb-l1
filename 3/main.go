/**
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с
использованием конкурентных вычислений.
*/

package main

import "fmt"

func sqrt(x int, c chan int) {
	c <- x * x
}

func main() {
	digits := [5]int{2, 4, 6, 8, 10}
	c := make(chan int)
	sum := 0

	for i := 0; i < len(digits); i++ {
		go sqrt(digits[i], c)
		sum += <-c
	}
	fmt.Println(sum)
}
