/**
Блокировка горутины с помощью контекста Deadline
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
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
	}

}
