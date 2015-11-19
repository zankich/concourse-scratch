package main

import "testing"

func TestReturnTrue(t *testing.T) {
	if ReturnTrue() != true {
		t.Error("Should have been true")
	}
}
