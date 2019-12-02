package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	ptr int
	seq []int
)

func process() bool {
	switch opCode := seq[ptr]; opCode {
	case 1:
		seq[seq[ptr+3]] = seq[seq[ptr+1]] + seq[seq[ptr+2]]
	case 2:
		seq[seq[ptr+3]] = seq[seq[ptr+1]] * seq[seq[ptr+2]]
	case 99:
		return true
	}
	ptr += 4
	return false
}

func readSequence(bytes []byte) []int {
	input := string(bytes)
	input = strings.TrimSuffix(input, "\n")
	var result []int
	parts := strings.Split(input, ",")
	for _, elem := range parts {
		number, _ := strconv.Atoi(elem)
		result = append(result, number)
	}
	return result
}

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	bytes, _ := ioutil.ReadFile(path)
	seq = readSequence(bytes)

	if path == "input.txt" {
		seq[1] = 12
		seq[2] = 2
	}

	for {
		if finished := process(); finished {
			break
		}
	}

	fmt.Println(seq)
}
