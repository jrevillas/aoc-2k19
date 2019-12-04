package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

type step struct {
	direction string
	length    int
}

func newStep(str string) step {
	length, _ := strconv.Atoi(str[1:])
	return step{
		direction: string(str[0]),
		length:    length,
	}
}

func parse(input []byte) [][]step {
	inputStr := string(input)
	inputStr = strings.TrimSuffix(inputStr, "\n")
	lines := strings.Split(inputStr, "\n")

	var result [][]step
	for _, line := range lines {
		stepsStr := strings.Split(line, ",")
		var steps []step
		for _, stepStr := range stepsStr {
			steps = append(steps, newStep(stepStr))
		}
		result = append(result, steps)
	}
	return result
}

type point struct {
	x, y int
}

func getPoints(steps []step) []point {
	var x, y int
	var path []point
	for _, step := range steps {
		for i := 0; i < step.length; i++ {
			switch step.direction {
			case up:
				y++
			case right:
				x++
			case down:
				y--
			case left:
				x--
			}
			path = append(path, point{x, y})
		}
	}
	return path
}

func nearest(m map[point]int) (point, int) {
	minDistance := -1
	var minPoint point
	for p, d := range m {
		if minDistance == -1 || d < minDistance {
			minDistance = d
			minPoint = p
		}
	}
	return minPoint, minDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isMemberTwice(s []point, e point) bool {
	var counter int
	for _, p := range s {
		if p.x == e.x && p.y == e.y {
			counter++
		}
	}
	return counter == 2
}

func wireDistance(ps []point, p point) int {
	var distance int
	for _, aux := range ps {
		distance++
		if aux.x == p.x && aux.y == p.y {
			break
		}
	}
	return distance
}

func wireDistanceSum(paths [][]step, point point) int {
	wireA := getPoints(paths[0])
	wireB := getPoints(paths[1])
	return wireDistance(wireA, point) + wireDistance(wireB, point)
}

func main() {
	path := "input.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	bytes, _ := ioutil.ReadFile(path)
	paths := parse(bytes)

	m := make(map[point]struct{})
	seen := make(map[point]int)
	for _, path := range paths {
		points := getPoints(path)
		for _, point := range points {
			if _, exist := m[point]; exist && !isMemberTwice(points, point) {
				seen[point] = wireDistanceSum(paths, point)
			} else {
				m[point] = struct{}{}
			}
		}
	}
	p, d := nearest(seen)
	fmt.Printf("nearest point: (%d,%d), distance: %d\n", p.x, p.y, d)
}
