package day4

import "fmt"

func fitsPasswordRequirementsPart1(input int) bool {
	prev := input
	sawDouble := false
	for i := 0; i < 6; i++ {
		current := input % 10
		if current > prev {
			return false
		}
		if prev == current {
			sawDouble = true
		}
		prev = current
		input = input / 10
	}
	return sawDouble
}

func fitsPasswordRequirementsPart2(input int) bool {
	prev := input

	counts := make(map[int]int)

	for i := 0; i < 6; i++ {
		current := input % 10
		if current > prev {
			return false
		}
		counts[current]++
		prev = current
		input = input / 10
	}

	for _, v := range counts {
		if v == 2 {
			return true
		}
	}

	return false
}

func PrintSolution() {
	passowordCount1 := 0
	passowordCount2 := 0

	for i := 123257; i <= 647015; i++ {
		if fitsPasswordRequirementsPart1(i) {
			passowordCount1++
		}
		if fitsPasswordRequirementsPart2(i) {
			passowordCount2++
		}
	}
	fmt.Println("There are ", passowordCount1, " different passwords in that range (Part 1)")
	fmt.Println("There are ", passowordCount2, " different passwords in that range (Part 2)")

}
