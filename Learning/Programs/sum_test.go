package main

import "testing"

func TestSum(t *testing.T) {
	x := 5
	y := 10
	want := Sum(x, y)
	get := 15
	if get != want {
		t.Errorf("get %d, want %d", get, want)
	}
}

func TestSub(t *testing.T) {
	x := 65
	y := 5
	get := Sub(x, y)
	want := 60
	if get != want {
		t.Errorf("get %d, want %d", get, want)
	}
}
