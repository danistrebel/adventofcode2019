package day4

import "testing"

func TestPasswordRequirementsPart1(t *testing.T) {
	testPasswordPart1(111111, true, t)
	testPasswordPart1(223450, false, t)
	testPasswordPart1(123789, false, t)
}

func testPasswordPart1(pass int, shouldPass bool, t *testing.T) {
	if fitsPasswordRequirementsPart1(pass) != shouldPass {
		t.Errorf("Password %d should have been checked with result %t", pass, shouldPass)
	}
}

func TestPasswordRequirementsPart2(t *testing.T) {
	testPasswordPart2(112233, true, t)
	testPasswordPart2(123444, false, t)
	testPasswordPart2(111122, true, t)
}

func testPasswordPart2(pass int, shouldPass bool, t *testing.T) {
	if fitsPasswordRequirementsPart2(pass) != shouldPass {
		t.Errorf("Password %d should have been checked with result %t", pass, shouldPass)
	}
}
