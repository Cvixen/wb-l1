/**
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import "fmt"

func uniqueValues(array []string) []string {
	set := make(map[string]struct{})
	var uniqueArray []string

	for _, value := range array {
		set[value] = struct{}{}
	}
	for k, _ := range set {
		uniqueArray = append(uniqueArray, k)
	}
	return uniqueArray
}

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(uniqueValues(array))

}
