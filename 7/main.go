/**
Реализовать конкурентную запись данных в map.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// Counters Структура с мапой и мьютексом
type Counters struct {
	mx sync.Mutex
	m  map[int]int
}

// NewCounters Инициализация структуры
func NewCounters() *Counters {
	return &Counters{
		m: make(map[int]int),
	}
}

// Store Потокобезопасня запись в мапу
func (c *Counters) Store(key int, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func main() {
	//Мапа из пакета синхронайз
	var mapBasicSync sync.Map
	//Собственная реализация потокобезопасной мапы
	mapSync := NewCounters()
	for i := 0; i < 100; i++ {
		go mapSync.Store(i, i)
		go mapBasicSync.Store(i, i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println(len(mapSync.m))
	fmt.Println(mapBasicSync)
}
