/**
Реализовать собственную функцию sleep.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func mySleep(n time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), n)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
func main() {
	start := time.Now()
	mySleep(3 * time.Second)
	duration := time.Since(start)
	fmt.Println(duration)
}
