package main

import "testing"

func TestAdd(t *testing.T) {
	if res := add(1, 2); res != 3 {
		t.Errorf("add(1, 2) = %d; want 3", res)
	}
}

func TestSub(t *testing.T) {
	if res := sub(1, 2); res != -1 {
		t.Errorf("sub(1, 2) = %d; want -1", res)
	}
}
