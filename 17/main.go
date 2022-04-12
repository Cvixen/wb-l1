/**
Реализовать бинарный поиск встроенными методами языка.
*/
package main

import (
	"fmt"
	"sort"
)

func SearchInIntSlice(haystack []int, needle int) int {
	sort.Ints(haystack)
	lowKey := 0
	highKey := len(haystack) - 1
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return -1 // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		mid := (lowKey + highKey) / 2 // середина
		if haystack[mid] == needle {
			return mid // нашли значение
		}
		if haystack[mid] < needle {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = mid - 1
		}
	}
	return -1
}

func main() {
	//Отсортированный 0, 1, 2, 5, 6, 7
	array := []int{5, 6, 7, 2, 1, 0}
	fmt.Println(SearchInIntSlice(array, 5))
}
