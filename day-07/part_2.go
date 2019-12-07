package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type amplifier struct {
	inputs   []int
	inputPtr int
	output   int
	ptr      int
	seq      []int
}

func newAmplifier(inputs, seq []int) *amplifier {
	a := amplifier{
		inputs: make([]int, len(inputs)),
		seq:    make([]int, len(seq)),
	}
	copy(a.inputs, inputs)
	copy(a.seq, seq)
	return &a
}

func (a *amplifier) getByMode(mode byte, x int) int {
	if mode == '0' {
		return a.seq[x]
	}
	return x
}

func (a *amplifier) readNextInput() int {
	a.inputPtr++
	return a.inputs[a.inputPtr-1]
}

func (a *amplifier) process() bool {
	opCode := fmt.Sprintf("%05d", a.seq[a.ptr])
	var result int
	switch op := opCode[len(opCode)-2:]; op {
	case "01":
		x := a.getByMode(opCode[2], a.seq[a.ptr+1])
		y := a.getByMode(opCode[1], a.seq[a.ptr+2])
		result = x + y
		a.seq[a.seq[a.ptr+3]] = result
		a.ptr += 4
	case "02":
		x := a.getByMode(opCode[2], a.seq[a.ptr+1])
		y := a.getByMode(opCode[1], a.seq[a.ptr+2])
		result = x * y
		a.seq[a.seq[a.ptr+3]] = result
		a.ptr += 4
	case "03":
		a.seq[a.seq[a.ptr+1]] = a.readNextInput()
		a.ptr += 2
	case "04":
		a.output = a.getByMode(opCode[2], a.seq[a.ptr+1])
		// fmt.Printf("#%d: %d\n", a.seq[a.ptr+1], a.output)
		a.ptr += 2
		return true
	case "05":
		if x := a.getByMode(opCode[2], a.seq[a.ptr+1]); x != 0 {
			a.ptr = a.getByMode(opCode[1], a.seq[a.ptr+2])
		} else {
			a.ptr += 3
		}
	case "06":
		if x := a.getByMode(opCode[2], a.seq[a.ptr+1]); x == 0 {
			a.ptr = a.getByMode(opCode[1], a.seq[a.ptr+2])
		} else {
			a.ptr += 3
		}
	case "07":
		x := a.getByMode(opCode[2], a.seq[a.ptr+1])
		y := a.getByMode(opCode[1], a.seq[a.ptr+2])
		if x < y {
			a.seq[a.seq[a.ptr+3]] = 1
		} else {
			a.seq[a.seq[a.ptr+3]] = 0
		}
		a.ptr += 4
	case "08":
		x := a.getByMode(opCode[2], a.seq[a.ptr+1])
		y := a.getByMode(opCode[1], a.seq[a.ptr+2])
		if x == y {
			a.seq[a.seq[a.ptr+3]] = 1
		} else {
			a.seq[a.seq[a.ptr+3]] = 0
		}
		a.ptr += 4
	case "99":
		return true
	}
	return false
}

func permutations(s []int) [][]int {
	var result [][]int

	var f func([]int, int)
	f = func(s []int, n int) {
		if n == 1 {
			tmp := make([]int, len(s))
			copy(tmp, s)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				f(s, n-1)
				if n%2 == 1 {
					tmp := s[i]
					s[i] = s[n-1]
					s[n-1] = tmp
				} else {
					tmp := s[0]
					s[0] = s[n-1]
					s[n-1] = tmp
				}
			}
		}
	}
	f(s, len(s))

	return result
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

func (a *amplifier) run() (int, bool) {
	for {
		if done := a.process(); done {
			break
		}
	}
	return a.output, a.seq[a.ptr] == 99
}

func createAmplifiers(pss, seq []int) []*amplifier {
	var as []*amplifier
	for _, ps := range pss {
		as = append(as, newAmplifier([]int{ps}, seq))
	}
	return as
}

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	bytes, _ := ioutil.ReadFile(path)
	seq := readSequence(bytes)

	var thrusterSignals []int
	for _, pss := range permutations([]int{5, 6, 7, 8, 9}) {
		as := createAmplifiers(pss, seq)
		var loops int
		var prevOutput int
		var prevOutputE int
		for i := 0; ; i++ {
			as[i-5*loops].inputs = append(as[i-5*loops].inputs, prevOutput)
			var halted bool
			prevOutput, halted = as[i-5*loops].run()
			if (i+1)%5 == 0 {
				prevOutputE = prevOutput
				loops++
			}
			if (i+1)%5 == 0 && halted {
				thrusterSignals = append(thrusterSignals, prevOutputE)
				break
			}
		}
	}

	sort.Ints(thrusterSignals)
	fmt.Printf("%d\n", thrusterSignals[len(thrusterSignals)-1])
}
