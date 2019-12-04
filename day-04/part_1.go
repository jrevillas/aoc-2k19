package main

import (
	"fmt"
	"strconv"
)

func ruleA(p int) bool {
	s := strconv.Itoa(p)
	var last int
	for _, rune := range s {
		digit := int(rune) - 48
		if digit < last {
			return false
		}
		last = digit
	}
	return true
}

func ruleB(p int) bool {
	s := strconv.Itoa(p)
	var prev rune
	for _, r := range s {
		if r == prev {
			return true
		}
		prev = r
	}
	return false
}

func main() {
	var passwords uint
	for p := 178416; p <= 676461; p++ {
		if ruleA(p) && ruleB(p) {
			passwords++
		}
	}
	fmt.Printf("%d\n", passwords)
}
