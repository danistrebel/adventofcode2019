package day5

import (
	"testing"
)

func TestOps(t *testing.T) {
	operation := parseOperation(1002)

	if operation.opcode != 2 {
		t.Error("Invalid Op Code:", operation.opcode, "expected:", 2)
	}
	if operation.modeParam1 != 0 {
		t.Error("Invalid Param1 Mode:", operation.modeParam1, "expected:", 0)
	}
	if operation.modeParam2 != 1 {
		t.Error("Invalid Param2 Mode:", operation.modeParam2, "expected:", 1)
	}
	if operation.modeParam3 != 0 {
		t.Error("Invalid Param3 Mode:", operation.modeParam3, "expected:", 0)
	}
}

func TestOpsExtended(t *testing.T) {
	if GravityAssistV2([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{8}) == 0 {
		t.Error("Expected output 0")
	}

}
