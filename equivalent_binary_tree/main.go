package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	length := 10
	ch1 := make(chan int, length)
	ch2 := make(chan int, length)
	values1 := sort.IntSlice([]int{})
	values2 := sort.IntSlice([]int{})

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < length; i++ {
		values1 = append(values1, <-ch1)
		values2 = append(values2, <-ch2)
	}

	go values1.Sort()
	go values2.Sort()

	time.Sleep(time.Millisecond * 1)

	for i := 0; i < length; i++ {
		if values1[i] != values2[i] {
			return false
		}
	}

	//fmt.Println(values1)
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	res := Same(t1, t2)

	fmt.Println(res)
}
