package client1

import "testing"

func TestReturnOne(t *testing.T) {
	if ReturnOne() != 1 {
		t.Error("Should have been 1")
	}
}
