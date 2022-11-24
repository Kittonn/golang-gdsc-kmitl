package main

import (
	"fmt"
)

func Question_1() {
	a := "BNK"
	b := 48
	c := true
	word := fmt.Sprintf("%v.%v.%v", a, b, c)
	fmt.Println(word)
}

func Question_2() {
	type kitton string
	var name kitton
	fmt.Printf("%T", name)
}

func main() {
	Question_1()
	fmt.Println("-----")
	Question_2()
}
