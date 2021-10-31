package day3

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// 1. index all points in wire 1
// 2. index points in wire 2

func nextPointInDirection(p Point, direction byte) Point {
	switch direction {
	case 'D':
		return Point{p.X, p.Y - 1}
	case 'U':
		return Point{p.X, p.Y + 1}
	case 'L':
		return Point{p.X - 1, p.Y}
	case 'R':
		return Point{p.X + 1, p.Y}
	}
	log.Fatal("unknown point")
	return Point{0, 0}
}

func manhattanDistance(p Point) int {
	distance := 0
	if p.X < 0 {
		distance += -p.X
	} else {
		distance += p.X
	}
	if p.Y < 0 {
		distance += -p.Y
	} else {
		distance += p.Y
	}
	return distance
}

func computeClosestIntersection(wire1 []string, wire2 []string) int {

	minDistance := MaxInt

	wire1Points := make(map[Point]bool)

	wire1current := Point{0, 0}
	for _, w := range wire1 {
		direction := w[0]
		distance, _ := strconv.Atoi(w[1:])
		for i := 0; i < distance; i++ {
			wire1current = nextPointInDirection(wire1current, direction)
			wire1Points[wire1current] = true
		}
	}

	wire2current := Point{0, 0}

	for _, w := range wire2 {
		direction := w[0]
		distance, _ := strconv.Atoi(w[1:])
		for i := 0; i < distance; i++ {
			wire2current = nextPointInDirection(wire2current, direction)
			if wire1Points[wire2current] {
				distanceFromOrigin := manhattanDistance(wire2current)
				if distanceFromOrigin < minDistance {
					minDistance = distanceFromOrigin
				}
			}
		}
	}

	return minDistance
}

func computeMinNumerOfSteps(wire1 []string, wire2 []string) int {

	minSteps := MaxInt

	wire1Points := make(map[Point]int)

	wire1current := Point{0, 0}
	wire1steps := 1

	for _, w := range wire1 {
		direction := w[0]
		distance, _ := strconv.Atoi(w[1:])
		for i := 0; i < distance; i++ {
			wire1current = nextPointInDirection(wire1current, direction)
			if _, ok := wire1Points[wire1current]; !ok {
				wire1Points[wire1current] = wire1steps
			}
			wire1steps++
		}
	}

	wire2current := Point{0, 0}
	wire2steps := 1
	for _, w := range wire2 {
		direction := w[0]
		distance, _ := strconv.Atoi(w[1:])
		for i := 0; i < distance; i++ {
			wire2current = nextPointInDirection(wire2current, direction)
			if wire1stepsmin, ok := wire1Points[wire2current]; ok {
				toalWireSteps := wire2steps + wire1stepsmin
				if toalWireSteps < minSteps {
					minSteps = toalWireSteps
				}
			}
			wire2steps++
		}
	}

	return minSteps
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day3.txt")
	distClosestIntersect := computeClosestIntersection(strings.Split(lines[0], ","), strings.Split(lines[1], ","))
	fmt.Println("Min distance intersection (Part 1): ", distClosestIntersect)
	minNumerOfSteps := computeMinNumerOfSteps(strings.Split(lines[0], ","), strings.Split(lines[1], ","))
	fmt.Println("Min wire steps intersection (Part 2): ", minNumerOfSteps)

}
