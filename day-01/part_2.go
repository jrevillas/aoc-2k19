package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func recursiveFuel(mass int) int {
	var total int
	for {
		fuel := mass/3 - 2
		if fuel <= 0 {
			break
		}
		total += fuel
		mass = fuel
	}
	return total
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
