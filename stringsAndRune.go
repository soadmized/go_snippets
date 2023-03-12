package main

import "fmt"

func main() {
	byteStr := "hello мир"
	runeStr := []rune(byteStr)

	for _, c := range byteStr {
		fmt.Println(c)
	}

	fmt.Println("--------")

	for _, r := range runeStr {
		fmt.Println(string(r))
	}

	fmt.Println("len byteStr - ", len(byteStr))
	fmt.Println("len runeStr - ", len(runeStr))

}
