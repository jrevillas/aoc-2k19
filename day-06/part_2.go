package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func basePlanetChain(m map[string]string, orbiter string) []string {
	var result []string
	for {
		if basePlanet, ok := m[orbiter]; ok {
			result = append(result, basePlanet)
			orbiter = basePlanet
		} else {
			break
		}
	}
	return result
}

func firstCommonAncestorDistance(a, b []string) int {
	for i := 0; ; i++ {
		if a[len(a)-i-1] != b[len(b)-i-1] {
			return len(a) + len(b) - 2*i
		}
	}
}

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	bytes, _ := ioutil.ReadFile(path)
	m := parseInput(bytes)

	chainA := basePlanetChain(m, "YOU")
	chainB := basePlanetChain(m, "SAN")
	d := firstCommonAncestorDistance(chainA, chainB)

	fmt.Printf("%d\n", d)
}
