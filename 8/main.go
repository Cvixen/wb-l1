/**
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

1 -> 00000001  -> смекщение на 1 бит -> 00000010
22 ->00010110           				00010110
 |   00010111                           00010110

00000010 - 1                           00000001
00010110 - 22						   00010110
00000010 - &						   00000000
00010100 - ^                           00010110
*/

package main

import "fmt"

func MoveBit(digit, moveBit, value int64) int64 {

	if value == 1 {
		return digit | int64(1<<moveBit)
	} else {
		return digit &^ int64(1<<moveBit)
	}
}

func main() {
	var digit, moveBit, value int64
	fmt.Println("Введите число, смещение на бит(от 0 до 63) и ноль или единицу")
	fmt.Scan(&digit, &moveBit, &value)

	fmt.Printf("До изменения: %064b\n", uint64(digit))
	newDigit := MoveBit(digit, moveBit, value)
	fmt.Printf("После: %064b\n", uint64(newDigit))

}
