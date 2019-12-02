package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fuel(mass int) int {
	f := func(mass, acc int) int {
		x := mass/3 - 2
		if x <= 0 {
			return acc
		}
		return f(x, acc+x)
	}
	return f(mass, 0)
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	content := string(bytes)

	var fuel int
	for _, module := range strings.Split(content, "\n") {
		if module == "" {
			continue
		}
		mass, _ := strconv.Atoi(module)
		fuel += recursiveFuel(mass)
	}

	fmt.Printf("fuel needed: %d\n", fuel)
}
