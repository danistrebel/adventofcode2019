package day2

import (
	"testing"
)

func TestGravityAssist(t *testing.T) {
	testcases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, testcase := range testcases {
		actualOutput := runGravityAssist(testcase.input)

		if !testEqual(actualOutput, testcase.output) {
			t.Errorf("Invalid Output for program: %d! Expected %d got %d!", testcase.input, testcase.output, actualOutput)
		}
	}
}

func testEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
