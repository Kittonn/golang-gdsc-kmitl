package exercise

import (
	"fmt"
)

func Exercise_1() {
	a := "BNK"
	b := 48
	c := true
	word := fmt.Sprintf("%v.%v.%v", a, b, c)
	fmt.Println(word)
	fmt.Println("-----")
	type kitton string
	var name kitton
	fmt.Printf("%T", name)
	fmt.Println()
	fmt.Println("-----")
}
