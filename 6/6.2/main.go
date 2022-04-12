/**
Блокировка горутины с помощью контекста Close
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 3 {
			cancel()
			break
		}
	}

}
