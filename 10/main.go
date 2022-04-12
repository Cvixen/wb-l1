/**
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
package main

import "fmt"

type Maps struct {
	//m2 используется только для значений (-10,0)
	m, m2 map[int][]float64
}

func temperature(arr []float64) Maps {
	m := make(map[int][]float64)
	m2 := make(map[int][]float64)
	var maps Maps
	for _, value := range arr {
		//проверка что ключ мапы == 0 и значение температуры меньше нуля
		if int(value/10)*10 == 0 && value < 0 {
			m2[0] = append(m2[0], value)
		} else {
			// для всех остальных значений добавляем в другую мапу
			m[int(value/10)*10] = append(m[int(value/10)*10], value)
		}
	}
	maps = Maps{
		m:  m,
		m2: m2,
	}
	return maps
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 4, -7, 0, -10}
	maps := temperature(arr)
	fmt.Println(maps.m)
	if maps.m2[0] != nil {
		fmt.Println(maps.m2)
	}
}
