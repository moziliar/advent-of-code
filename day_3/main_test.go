package main

import "testing"

func TestWithin(t *testing.T) {
	got := within(5, 3, 9)
	if got != true {
		t.Error("within(5, 3, 9) = ", got, "want true")
	}
	got = within(5,8,1)
	if got != true {
		t.Error("within(5, 8, 1) = ", got, "want true")
	}
}

func TestParallel(t *testing.T) {
	got := parallel(edge{point{0, 0}, point{1000, 0}}, edge{point{518, 1023}, point{518, 669}})
	if got != false {
		t.Error("parallel(e1, e2 = ", got, "want false")
	}
}
