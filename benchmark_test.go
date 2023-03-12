package main

import (
	"math/rand"
	"testing"
)

func BenchmarkX(b *testing.B) {
	type digit int
	const (
		zero = iota
		one
		two
		three
		four
		five
		six
		seven
		eight
		nine
	)

	switchCase := func(d digit) string {
		switch d {
		case zero:
			return "zero"
		case one:
			return "one"
		case two:
			return "two"
		case three:
			return "three"
		case four:
			return "four"
		case five:
			return "five"
		case six:
			return "six"
		case seven:
			return "seven"
		case eight:
			return "eight"
		case nine:
			return "nine"
		}
		return ""
	}

	m := map[digit]string{
		zero:  "zero",
		one:   "one",
		two:   "two",
		three: "three",
		four:  "four",
		five:  "five",
		six:   "six",
		seven: "seven",
		eight: "eight",
		nine:  "nine",
	}

	s := []struct {
		digit
		name string
	}{
		{zero, "zero"},
		{one, "one"},
		{two, "two"},
		{three, "three"},
		{four, "four"},
		{five, "five"},
		{six, "six"},
		{seven, "seven"},
		{eight, "eight"},
		{nine, "nine"},
	}

	byMap := func(d digit) string {
		return m[d]
	}

	bySlice := func(d digit) string {
		for i := 0; i < 10; i++ {
			if s[i].digit == d {
				return s[i].name
			}
		}

		return ""
	}

	b.Run("switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switchCase(digit(rand.Intn(10)))
		}
	})

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			byMap(digit(rand.Intn(10)))
		}
	})

	b.Run("slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bySlice(digit(rand.Intn(10)))
		}
	})
}
