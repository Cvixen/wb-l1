/**
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func worker(id int, input <-chan int) {
	for i := range input {
		fmt.Printf("[%d Worker] Worker started  job  %d\n", id, i)
	}

}
func main() {
	var countWorkers int

	_, err := fmt.Scan(&countWorkers)
	if err != nil {
		log.Fatalln(err)
	}

	input := make(chan int)
	for i := 0; i < countWorkers; i++ {
		go worker(i, input)
	}

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)

	i := 0

	func() {
		for {
			time.Sleep(500 * time.Millisecond)
			select {
			case <-signalChanel:
				return
			default:
				input <- i
				i++
			}
		}
	}()
	close(input)
	close(signalChanel)
}
