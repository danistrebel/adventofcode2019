package day1

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
)

func fuelCalc(weight int) int {
	return (weight / 3) - 2
}

func fuelForFuelCalc(weight int) int {
	fuel := fuelCalc(weight)
	if fuel > 0 {
		return fuel + fuelForFuelCalc(fuel)
	}
	return 0
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day1.txt")

	total_fuel := 0
	total_fuel_for_fuel := 0

	for _, line := range lines {
		w, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		total_fuel += fuelCalc(w)
		total_fuel_for_fuel += fuelForFuelCalc(w)
	}

	fmt.Println("Total Fuel (Part 1):", total_fuel)
	fmt.Println("Total Fuel For Fuel (Part 2):", total_fuel_for_fuel)
}
