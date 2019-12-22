package main

import "testing"

func TestAccept(t *testing.T) {
	type test struct {
		input int
		ans   bool
	}
	tests := []test{
		{111111, true},
		{123450, false},
		{567899, true},
	}
	for _, ts := range tests {
		got := accept(ts.input)
		if got != ts.ans {
			t.Error("accept(", ts.input, ") = ", got, "want", ts.ans)
		}
	}
}
