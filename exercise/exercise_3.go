package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 5)
	for i := 1; i < 10; i += 2 {
		s = append(s, i)
	}
	s = append([]int{0}, s...)
	fmt.Println(s)

	fmt.Println("-----")

	removeItem := s[:len(s)-1]
	fmt.Println(removeItem)

	fmt.Println("-----")

	data := make(map[string][]string)
	data["Alex"] = []string{"icecream", "prata"}
	data["Bob"] = []string{"pizza", "taco"}
	for key, value := range data {
		fmt.Println(key, value)
	}
}
