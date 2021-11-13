package day5

import (
	"adventofcode/utils"
	"fmt"
)

type Instruction struct {
	opcode     int
	modeParam1 int
	modeParam2 int
	modeParam3 int
}

func parseOperation(i int) Instruction {
	return Instruction{
		(i % 100), (i % 1000 / 100), (i % 10000 / 1000), (i / 10000),
	}
}

func paramParser(program []int, i int, mode int) int {
	if mode == 0 {
		return program[program[i]]
	} else if mode == 1 {
		return program[i]
	} else {
		panic("invalid mode")
	}
}

func runGravityAssistV2(program []int, input int) int {
	i := 0
	var diagnosticOut int
	for i < len(program) {
		op := parseOperation(program[i])
		if op.opcode == 99 {
			break
		} else {
			switch op.opcode {
			case 1:
				program[program[i+3]] = paramParser(program, i+1, op.modeParam1) + paramParser(program, i+2, op.modeParam2)
				i += 4
			case 2:
				program[program[i+3]] = paramParser(program, i+1, op.modeParam1) * paramParser(program, i+2, op.modeParam2)
				i += 4
			case 3:
				program[program[i+1]] = input
				i += 2
			case 4:
				if program[i+2] == 99 {
					diagnosticOut = program[program[i+1]]
				}
				i += 2
			case 5:
				if paramParser(program, i+1, op.modeParam1) != 0 {
					i = paramParser(program, i+2, op.modeParam2)
				} else {
					i += 3
				}
			case 6:
				if paramParser(program, i+1, op.modeParam1) == 0 {
					i = paramParser(program, i+2, op.modeParam2)
				} else {
					i += 3
				}
			case 7:
				if paramParser(program, i+1, op.modeParam1) < paramParser(program, i+2, op.modeParam2) {
					program[program[i+3]] = 1
				} else {
					program[program[i+3]] = 0
				}
				i += 4
			case 8:
				if paramParser(program, i+1, op.modeParam1) == paramParser(program, i+2, op.modeParam2) {
					program[program[i+3]] = 1
				} else {
					program[program[i+3]] = 0
				}
				i += 4
			}
		}
	}
	return diagnosticOut
}

func PrintSolution() {
	output := runGravityAssistV2(utils.ParseIntArrays("./inputs/day5.txt")[0], 1)
	fmt.Println("Program Output (Part 1):", output)
	output2 := runGravityAssistV2(utils.ParseIntArrays("./inputs/day5.txt")[0], 5)
	fmt.Println("Program Output (Part 2):", output2)

}
