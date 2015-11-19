package client2

import "testing"

func TestReturnTwo(t *testing.T) {
	if ReturnTwo() != 2 {
		t.Error("Should have been 2")
	}
}
