package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func indirect(m map[string]string, orbiter string) int {
	if basePlanet, ok := m[orbiter]; ok {
		return 1 + indirect(m, basePlanet)
	}
	return 0
}

func parseInput(bytes []byte) map[string]string {
	input := string(bytes)
	input = strings.TrimSuffix(input, "\n")
	result := make(map[string]string)
	parts := strings.Split(input, "\n")
	for _, elem := range parts {
		elemParts := strings.Split(elem, ")")
		result[elemParts[1]] = elemParts[0]
	}
	return result
}

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	bytes, _ := ioutil.ReadFile(path)
	m := parseInput(bytes)

	var total int
	for orbiter := range m {
		total += indirect(m, orbiter)
	}

	fmt.Printf("%d\n", total)
}
