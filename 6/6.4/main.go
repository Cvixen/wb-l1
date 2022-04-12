/**
Блокировка горутины с помощью останавливающего канала
*/
package main

import (
	"fmt"
	"time"
)

func gen(stop chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-stop:
				close(ch)
				return
			case ch <- n:
				time.Sleep(time.Second)
				n++
			}
		}
	}()
	return ch
}

func main() {
	stop := make(chan bool)

	for n := range gen(stop) {
		fmt.Println(n)
		if n == 3 {
			stop <- true
			break
		}
	}

}
