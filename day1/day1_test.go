package day1

import "testing"

func TestFuelCalc(t *testing.T) {
	testcases := []struct {
		weight int
		fuel   int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, testcase := range testcases {
		fuelCalculated := fuelCalc(testcase.weight)
		if fuelCalculated != testcase.fuel {
			t.Errorf("Invalid Fuel for weight: %d! Expected %d got %d!", testcase.weight, testcase.fuel, fuelCalculated)
		}
	}
}

func TestFuelForFuelCalc(t *testing.T) {
	testcases := []struct {
		weight int
		fuel   int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, testcase := range testcases {
		fuelCalculated := fuelForFuelCalc(testcase.weight)
		if fuelCalculated != testcase.fuel {
			t.Errorf("Invalid Fuel for weight: %d! Expected %d got %d!", testcase.weight, testcase.fuel, fuelCalculated)
		}
	}
}
