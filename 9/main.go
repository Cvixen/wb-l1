/**
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/
package main

import (
	"fmt"
)

func main() {
	digits := [5]int{2, 4, 6, 8, 10}
	input := make(chan int)
	output := make(chan int)
	wait := make(chan struct{})
	// Из output отправляем значения в stdout
	go func() {
		for j := range output {
			fmt.Println(j)
		}
		wait <- struct{}{}
	}()
	// Из канала input отправляем значения в output
	go func() {
		for i := range input {
			output <- i * i
		}
		close(output)
	}()
	// Отправляем значения в канал input
	for i := 0; i < len(digits); i++ {
		input <- digits[i]
	}
	close(input)
	<-wait
}
