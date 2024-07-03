package main

import (
	"fmt"
	"math/rand"
)

const epsilon = 0.1

func main() {
	explore := make([]bool, 0)
	exploit := make([]bool, 0)

	for i := 0; i < 100; i++ {
		prob := Probability()
		if prob {
			exploit = append(exploit, prob)
		} else {
			explore = append(explore, prob)
		}
	}

	fmt.Println("EXPLOIT = ", len(exploit))
	fmt.Println("EXPLORE = ", len(explore))

	m := map[string]int{
		"banner1": 1,
		"banner2": 2,
		"banner3": 2,
		"banner4": 3,
	}

	banner := maxValueInMap(m)

	fmt.Println(banner)
}

func Probability() bool {
	prob := rand.Float64()

	if prob <= (1 - epsilon) {
		return true // exploit
	} else {
		return false // explore
	}
}

func maxValueInMap(m map[string]int) string {
	maxV := 0
	maxK := ""

	for k, v := range m {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}

	return maxK
}
