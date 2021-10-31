package day3

import "testing"

func TestManhattanDistance(t *testing.T) {
	testcases := []struct {
		point    Point
		distance int
	}{
		{Point{0, 1}, 1},
		{Point{3, 3}, 6},
		{Point{-10, 1}, 11},
		{Point{-10, -33}, 43},
	}

	for _, testcase := range testcases {
		computeddistance := manhattanDistance(testcase.point)
		if computeddistance != testcase.distance {
			t.Errorf("Invalid distance for point %v got %d expected %d", testcase.point, computeddistance, testcase.distance)
		}
	}
}

func TestWireIntersect(t *testing.T) {
	testcases := []struct {
		input1 []string
		input2 []string
		output int
	}{
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			159,
		},
		{
			[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			135,
		},
	}

	for _, testcase := range testcases {
		closestintersection := computeClosestIntersection(testcase.input1, testcase.input2)

		if closestintersection != testcase.output {
			t.Errorf("Invalid distance, calculated %d, expected %d", closestintersection, testcase.output)
		}
	}
}
