/**
Реализовать паттерн «адаптер» на любом примере
*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

type client struct {
}

func (client *client) calculating(calculator calculate) {
	calculator.squaringDigits()
}

type calculate interface {
	squaringDigits()
}

type digits struct {
	x int
}

func (digits *digits) squaringDigits() {
	fmt.Printf("Quadratic number is %d\n", digits.x*2)
}

type stringDigit struct {
	y string
}

func (stringDigit *stringDigit) squaringString() {
	fmt.Printf("Quadratic number where digit like string is %s * 2\n", stringDigit.y)
}

type stringDigitAdapter struct {
	stringDigit *stringDigit
}

func (stringAdapter *stringDigitAdapter) squaringDigits() {
	digit, err := strconv.Atoi(stringAdapter.stringDigit.y)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Quadratic number from string is %d\n", digit*2)

}

func main() {
	digit := &digits{x: 3}
	client := client{}
	//клиент использует калькулятор с числами
	client.calculating(digit)

	stringDigits := stringDigit{y: "4"}
	//у строки есть способ считать в квадрате
	stringDigits.squaringString()
	//создание адаптера для строки в число
	adapter := &stringDigitAdapter{&stringDigits}
	//клиент использует калькулятор с адаптером
	client.calculating(adapter)

}
