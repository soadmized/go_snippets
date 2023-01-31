package main

import "fmt"

var f = func(x int) {}

func Foo() {
	f := func(x int) {
		if x >= 0 {
			fmt.Println(x)
			f(x - 1)
		}
	}
	f(2)
}

func Bar() {
	f = func(x int) {
		if x >= 0 {
			fmt.Println(x)
			f(x - 1)
		}
	}
	f(2)
}

func main() {
	Foo()
	fmt.Println(" | ")
	Bar()
}
