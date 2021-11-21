package day7

import (
	"adventofcode/day5"
	"adventofcode/utils"
	"fmt"
)

func runSequence(program []int, config []int) int {
	thrust := 0

	for configIndex := 0; configIndex < len(config); configIndex++ {
		thrust = day5.GravityAssistV2(program, []int{config[configIndex], thrust})
	}

	return thrust
}

func findMaxThrust(program []int) int {
	config := []int{0, 1, 2, 3, 4}

	return findMaxThrustFromFixed(program, config, 0)
}

func findMaxThrustFromFixed(program []int, config []int, positionsFixed int) int {

	if len(config) == positionsFixed {
		return runSequence(program, config)
	}

	max := 0

	for i := positionsFixed; i < len(config); i++ {
		configCopy := append([]int{}, config...)
		configCopy[positionsFixed], configCopy[i] = configCopy[i], configCopy[positionsFixed]

		thrustForConfig := findMaxThrustFromFixed(program, configCopy, positionsFixed+1)

		if thrustForConfig > max {
			max = thrustForConfig
		}
	}

	return max
}

func PrintSolution() {
	maxThrust := findMaxThrust(utils.ParseIntArrays("./inputs/day7.txt")[0])
	fmt.Println("Max Thrust (Part 1):", maxThrust)
}
