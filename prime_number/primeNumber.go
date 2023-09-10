package main

import (
	"fmt"
	"math"
)

func IsPrime(x int) bool {
	y := int(math.Sqrt(float64(x)))

	for i := 2; i < (y + 1); i++ {
		if (x % i) == 0 {
			return false
		}
	}

	return true
}

func main() {
	x := 5
	fmt.Printf("%d is prime: %t", x, IsPrime(x))
}
