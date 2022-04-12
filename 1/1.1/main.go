/**
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

type Human struct {
	FirstName, SecondName string
}

func (person *Human) say(words string) {
	fmt.Println(words)
}

type Action struct {
	Human Human
}

func main() {
	action := Action{
		Human{
			FirstName:  "Eugene",
			SecondName: "Blednykh",
		}}
	action.Human.say("Haha, classic")
}
