package day2

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

func runGravityAssist(program []int) []int {
	i := 0
	for i < len(program) {
		if program[i] == 99 {
			break
		} else if program[i] == 1 {
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		} else if program[i] == 2 {
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		} else {
			log.Fatal("Unexpected operand", program[i])
		}
		i += 4
	}
	return program
}

func freshInput(first int, second int) []int {
	intIntputs := utils.ParseIntArrays("./inputs/day2.txt")
	programInput := intIntputs[0]
	programInput[1] = first
	programInput[2] = second
	return programInput
}

func PrintSolution() {
	programInput := freshInput(12, 2)
	output := runGravityAssist(programInput)
	fmt.Println("Program Output (Part 1):", output[0])

	target := 19690720
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			if runGravityAssist(freshInput(n, v))[0] == target {
				fmt.Println("Solution for input", (100*n + v))
				break
			}
		}
	}

}
