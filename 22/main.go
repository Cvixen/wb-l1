/**
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(2000002)
	y := big.NewInt(2000000)

	fmt.Println(new(big.Int).Add(x, y)) //сложение
	fmt.Println(new(big.Int).Sub(x, y)) //вычитание
	fmt.Println(new(big.Int).Mul(x, y)) //умножение
	fmt.Println(new(big.Int).Div(x, y)) //деление
}
