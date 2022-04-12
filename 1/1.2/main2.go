package main

import "fmt"

type ActionHuman interface {
	say(words string)
}

type Human struct {
	FirstName, SecondName string
}

func (person *Human) say(words string) {
	fmt.Println("Human say" + words)
}

type Action struct{}

func (action *Action) say(words string) {
	fmt.Println("Action say: " + words)
}
func main() {
	action := new(Action)
	action.say("Haha, classic")
}
