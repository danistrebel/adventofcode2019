package day7

import "testing"

func TestSequenceThrust(t *testing.T) {
	testcases := []struct {
		input  []int
		config []int
		result int
	}{
		{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{4, 3, 2, 1, 0}, 43210},
		{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, []int{0, 1, 2, 3, 4}, 54321},
		{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []int{1, 0, 4, 3, 2}, 65210},
	}

	for _, test := range testcases {
		calculated := runSequence(test.input, test.config)
		if calculated != test.result {
			t.Error("Wrong thrust result: Got", calculated, "expected:", test.result)
		}
	}
}

func TestMaxThrust(t *testing.T) {
	testcases := []struct {
		input  []int
		result int
	}{
		{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
		{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
		{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, 65210},
	}

	for _, test := range testcases {
		calculated := findMaxThrust(test.input)
		if calculated != test.result {
			t.Error("Wrong thrust result: Got", calculated, "expected:", test.result)
		}
	}
}
