/**
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

func typeDeterminant(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
}
func main() {
	integer := 20
	str := "lol"
	bl := true
	channel := make(chan int)

	typeDeterminant(integer)
	typeDeterminant(str)
	typeDeterminant(bl)
	typeDeterminant(channel)
}
