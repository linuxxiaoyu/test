package main

import "testing"

func TestSub(t *testing.T) {
	if sub(1, 2) != -1 {
		t.Error("1-2=-1")
	}
}
