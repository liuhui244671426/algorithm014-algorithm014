package Week_02

import (
	"fmt"
)

func fizzBuzz(n int) []string {

	var m []string
	var f string
	for i := 1; i <= n; i++ {
		//fmt.Println(i%3, i%5)
		if (i%3 == 0) && (i%5 == 0) {
			f = "FizzBuzz"
		} else if i%5 == 0 {
			f = "Buzz"
		} else if i%3 == 0 {
			f = "Fizz"
		} else {
			f = fmt.Sprintf("%d", i)
		}
		fmt.Println(f)
		m = append(m, f)
	}
	return m
}

func FizzBuzz(n int) []string {
	return fizzBuzz(n)
}
