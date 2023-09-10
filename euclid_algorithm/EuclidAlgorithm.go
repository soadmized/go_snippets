package main

import "fmt"

func EuclidAlgorithm[T int | int64 | int32 | int8 | int16](x, y T) T {
	if y == 0 {
		return x
	}

	for y != 0 {
		x, y = y, (x % y)
	}

	return x
}

func main() {
	x, y := 12, 20
	fmt.Println(EuclidAlgorithm(x, y))
}
