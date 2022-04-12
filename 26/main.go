/**
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func uniqueSym(str string) bool {
	str = strings.ToLower(str)
	cashSym := make(map[rune]bool)

	for _, val := range str {
		_, ok := cashSym[val]
		if ok {
			return false
		} else {
			cashSym[val] = true
		}
	}
	return true
}
func main() {
	fmt.Println(uniqueSym("abcd"))
	fmt.Println(uniqueSym("abCdefAaf"))
	fmt.Println(uniqueSym("aabcd"))
	fmt.Println(uniqueSym("женя"))
}
