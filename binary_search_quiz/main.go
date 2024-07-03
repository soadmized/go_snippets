package main

import (
	"fmt"
	"strconv"
)

func main() {
	greeting()

	lower := 0
	upper := initialInput()
	res := search(lower, upper)

	fmt.Println(res)
}

func initialInput() int {
	var input string

	fmt.Println("Загадайте число :) И введите верхнюю границу для поиска (целое число):")
	fmt.Scanf("%s\\n", &input)

	limit := validateIntInput(input)

	return limit
}

func greeting() {
	greet := "================================================\n|--\\  |  |\\  |   /-\\   |--\\   |   |\n|_ /  |  | \\ |  |   |  |   )  \\   /\n|  \\  |  |  \\|  |___|  |--/    \\ /\n|__/  |  |   |  |   |  |  \\     |\n\n\t     ----  |---   /-\\   |--\\   |---   |   |\n\t     |     |     |   |  |   )  |      |   |\n\t     |---  |---  |___|  |--/   |      |---|\n\t        |  |     |   |  |  \\   |      |   |\n\t     ----  ----  -   -  -   -  ----   -   -\n================================================"
	fmt.Println(greet)
}

func search(lower, upper int) int {
	if lower == upper { // patch for mistake user answers
		lower = -1
		upper = +1
	}

	try := (lower + upper) / 2
	answer := askTry(try)

	switch answer {
	case 0:
		return try
	case 1: // lower than try
		return search(lower, try-1)
	case 2: // higher than try
		return search(try+1, upper)
	}

	return 0
}

func askTry(try int) int {
	var input string

	fmt.Printf("Загаданное число больше или меньше %d?\n(0 - это и есть мое число, 1 - меньше, 2 - больше):\n", try)
	fmt.Scanf("%s\\n", &input)

	answer := validateIntInput(input)

	return answer
}

func validateIntInput(input string) int {
	var res int
	var err error

	for {
		res, err = strconv.Atoi(input)
		if err != nil {
			fmt.Printf("%s - некорректное число. \nВведите заново:\n", input)
			fmt.Scanf("%s\\n", &input)
		} else {
			return res
		}
	}
}
