package main

import "testing"

func TestAdd(t *testing.T) {
	if add(1, 1) != 2 {
		t.Error("1+1=2")
	}
}
