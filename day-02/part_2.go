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

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			seq = readSequence(bytes)
			ptr = 0
			seq[1] = noun
			seq[2] = verb
			for {
				if finished := process(); finished {
					if seq[0] == 19690720 {
						fmt.Printf("%d\n", 100*noun+verb)
						return
					}
					break
				}
			}
		}
	}
}
