package main

import (
	"fmt"
	"time"
)

func Question_1() {
	for i := 6; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func Question_2() {
	j := 10
	for j > 0 {
		if j == 3 {
			break
		}
		if j%2 != 0 {
			fmt.Println(j)
		}
		j--
	}
}

func Question_3() {
	day := time.Now().Weekday()
	switch day {
	case time.Saturday:
		word := fmt.Sprintf("%v is weekend", day)
		fmt.Println(word)
	case time.Sunday:
		word := fmt.Sprintf("%v is weekend", day)
		fmt.Println(word)
	default:
		word := fmt.Sprintf("%v is weekday", day)
		fmt.Println(word)
	}
}

func main() {
	Question_1()
	fmt.Println("-----")
	Question_2()
	fmt.Println("-----")
	Question_3()
}
