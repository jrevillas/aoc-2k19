package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	input = 5
	ptr   int
	seq   []int
)

func process() bool {
	switch opCode := seq[ptr]; opCode {
	case 99:
		return true
	default:
		handleComplexOpCode()
	}
	return false
}

func getByMode(mode byte, x int) int {
	if mode == '0' {
		return seq[x]
	}
	return x
}

func handleComplexOpCode() {
	opCode := fmt.Sprintf("%05d", seq[ptr])
	var result int
	switch op := opCode[len(opCode)-2:]; op {
	case "01":
		x := getByMode(opCode[2], seq[ptr+1])
		y := getByMode(opCode[1], seq[ptr+2])
		result = x + y
		seq[seq[ptr+3]] = result
		ptr += 4
	case "02":
		x := getByMode(opCode[2], seq[ptr+1])
		y := getByMode(opCode[1], seq[ptr+2])
		result = x * y
		seq[seq[ptr+3]] = result
		ptr += 4
	case "03":
		seq[seq[ptr+1]] = input
		ptr += 2
	case "04":
		x := getByMode(opCode[2], seq[ptr+1])
		fmt.Printf("#%d: %d\n", seq[ptr+1], x)
		ptr += 2
	case "05":
		if x := getByMode(opCode[2], seq[ptr+1]); x != 0 {
			ptr = getByMode(opCode[1], seq[ptr+2])
		} else {
			ptr += 3
		}
	case "06":
		if x := getByMode(opCode[2], seq[ptr+1]); x == 0 {
			ptr = getByMode(opCode[1], seq[ptr+2])
		} else {
			ptr += 3
		}
	case "07":
		x := getByMode(opCode[2], seq[ptr+1])
		y := getByMode(opCode[1], seq[ptr+2])
		if x < y {
			seq[seq[ptr+3]] = 1
		} else {
			seq[seq[ptr+3]] = 0
		}
		ptr += 4
	case "08":
		x := getByMode(opCode[2], seq[ptr+1])
		y := getByMode(opCode[1], seq[ptr+2])
		if x == y {
			seq[seq[ptr+3]] = 1
		} else {
			seq[seq[ptr+3]] = 0
		}
		ptr += 4
	}
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

	for {
		if finished := process(); finished {
			break
		}
	}
}
