package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	content := string(bytes)

	var fuel int
	for _, module := range strings.Split(content, "\n") {
		if module == "" {
			continue
		}
		mass, _ := strconv.Atoi(module)
		fuel += mass/3 - 2
	}

	fmt.Printf("fuel needed: %d\n", fuel)
}
