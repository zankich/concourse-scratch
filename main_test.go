package main

import "testing"

func TestReturnThree(t *testing.T) {
	if ReturnThree() != 3 {
		t.Error("Should have been 3")
	}
}
