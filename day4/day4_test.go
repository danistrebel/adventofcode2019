package day4

import "testing"

func TestPasswordRequirements(t *testing.T) {
	testPassword(111111, true, t)
	testPassword(223450, false, t)
	testPassword(123789, false, t)
}

func testPassword(pass int, shouldPass bool, t *testing.T) {
	if fitsPasswordRequirements(pass) != shouldPass {
		t.Errorf("Password %d should have been checked with result %t", pass, shouldPass)
	}
}
