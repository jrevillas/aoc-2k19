package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func parseLayers(bs []byte, size int) [][]int {
	bs = bytes.TrimSuffix(bs, []byte("\n"))

	var input []int
	for _, b := range bs {
		number, _ := strconv.Atoi(string(b))
		input = append(input, number)
	}

	var layers [][]int
	for i := 0; i < len(input)/size; i++ {
		layer := input[size*i : size*(i+1)]
		layers = append(layers, layer)
	}

	return layers
}

func count(layer []int, x int) int {
	var counter int
	for _, pixel := range layer {
		if pixel == x {
			counter++
		}
	}
	return counter
}

func minPosition(s []int) int {
	min := s[0]
	minPos := 0
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
			minPos = i
		}
	}
	return minPos
}

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	bytes, _ := ioutil.ReadFile(path)
	size := 25 * 6
	layers := parseLayers(bytes, size)

	var zerosInLayers []int
	for _, layer := range layers {
		zerosInLayers = append(zerosInLayers, count(layer, 0))
	}

	layer := layers[minPosition(zerosInLayers)]
	fmt.Printf("%d\n", count(layer, 1)*count(layer, 2))
}
