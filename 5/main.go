/**
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	var count time.Duration
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	go func() {
		i := 0
		for {
			c <- i
			i++
			time.Sleep(time.Second)
		}
	}()

	go func(int) {
		for {
			fmt.Printf("Data from channel: %d\n", <-c)
		}
	}(<-c)

	time.Sleep(time.Second * count)
}
