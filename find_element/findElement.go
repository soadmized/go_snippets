package main

import "fmt"

// FindElement returns the index of x in s, or -1 if not found.
func FindElement[T comparable](haystack []T, needle T) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func main() {
	// FindElement works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(FindElement(si, 15))

	// FindElement also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(FindElement(ss, "hello"))
}
