package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	height = 6
	width  = 25
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

func getPixelInLayer(layer []int, x, y int) int {
	return layer[y*width+x]
}

func getPixel(layers [][]int, x, y int) string {
	result := 2
	for i := len(layers) - 1; i >= 0; i-- {
		if pixel := getPixelInLayer(layers[i], x, y); pixel != 2 {
			result = pixel
		}
	}
	if result == 1 {
		return "#"
	}
	return " "
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
	layers := parseLayers(bytes, height*width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%s", getPixel(layers, x, y))
		}
		fmt.Print("\n")
	}
}
