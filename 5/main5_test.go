package main

import (
	"testing"
)

func containsSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	set := make(map[int]int)
	for _, v := range a {
		set[v]++
	}

	for _, v := range b {
		set[v]--
		if set[v] < 0 {
			return false
		}
	}

	return true
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name        string
		a           []int
		b           []int
		expectedOK  bool
		expectedRes []int
	}{
		{
			"intersection",
			[]int{65, 3, 58, 678, 64},
			[]int{64, 2, 3, 43},
			true,
			[]int{64, 3},
		},
		{
			"no intersection",
			[]int{1, 2},
			[]int{3, 4},
			false,
			[]int{},
		},
		{
			"one empty",
			[]int{},
			[]int{1, 2},
			false,
			[]int{},
		},
		{
			"duplicate",
			[]int{1, 2, 2, 3},
			[]int{2, 2, 4},
			true,
			[]int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, res := Intersection(tt.a, tt.b)

			if ok != tt.expectedOK {
				t.Fatalf("got ok=%v, want %v", ok, tt.expectedOK)
			}

			if !containsSameElements(res, tt.expectedRes) {
				t.Fatalf("got %v, want %v", res, tt.expectedRes)
			}
		})
	}
}
