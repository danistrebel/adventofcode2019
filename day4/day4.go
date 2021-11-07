package day4

import "fmt"

func fitsPasswordRequirements(input int) bool {
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

func PrintSolution() {
	passowordCount := 0
	for i := 123257; i <= 647015; i++ {
		if fitsPasswordRequirements(i) {
			passowordCount++
		}
	}
	fmt.Println("There are ", passowordCount, " different passwords in that range (Part 1)")
}
