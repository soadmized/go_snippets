package main

import "fmt"

func main() {
	number := 15
	fizzBuzz(number)
}

func fizzBuzz[T Numbers](number T) {
	var i T
	for i = 1; i <= number; i++ {
		if ((i % 5) == 0) && ((i % 3) == 0) {
			fmt.Println("FIZZBUZZ", i)
		} else if (i % 3) == 0 {
			fmt.Println("FIZZ", i)
		} else if (i % 5) == 0 {
			fmt.Println("BUZZ", i)
		} else {
			fmt.Println(i)
		}
	}
}

type Numbers interface {
	int | int64 | int32 |
	int8 | int16
}
