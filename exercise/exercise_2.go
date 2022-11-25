package exercise

import (
	"fmt"
	"time"
)

func Exercise_2() {
	for i := 6; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("-----")
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
	fmt.Println("-----")
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
	fmt.Println("-----")
}
