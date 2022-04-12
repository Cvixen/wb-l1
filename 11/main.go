/**
Реализовать пересечение двух неупорядоченных множеств.
*/
package main

import "fmt"

func IntersectionOfSet(array1, array2 []int) []int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	interSet := make(map[int]int)
	//Создаем сет от первого массива
	for _, value := range array1 {
		set1[value] = struct{}{}
	}
	//Создаем сет от второго массива
	for _, value := range array2 {
		set2[value] = struct{}{}
	}
	//Вставляем в ключи значение +1. Там где число будет 2, значит есть пересечение
	for key, _ := range set1 {
		interSet[key] += 1
	}
	for key, _ := range set2 {
		interSet[key] += 1
	}
	//Преобразуем сет в слайс
	var array []int
	for k, v := range interSet {
		if v == 2 {
			array = append(array, k)
		}
	}
	return array
}

func main() {
	array1 := []int{1, 2, 5, 7, 3, 2, 7, 1, 6, 56, 23}
	array2 := []int{1, 2, 5, 7, 34, 2, 4, 213, 8, 56, 23}
	fmt.Println(IntersectionOfSet(array1, array2))

}
