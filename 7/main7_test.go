package main

import (
	"testing"
)

func collect[T any](ch <-chan T) []T {
	var res []T
	for v := range ch {
		res = append(res, v)
	}
	return res
}

func sameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[int]int)
	for _, v := range a {
		count[v]++
	}
	for _, v := range b {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  []int
	}{
		{
			name: "2 chs",
			input: [][]int{
				{1, 2, 3},
				{4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "1 ch",
			input: [][]int{
				{10, 20},
			},
			want: []int{10, 20},
		},
		{
			name: "empty chs",
			input: [][]int{
				{},
				{},
			},
			want: []int{},
		},
		{
			name:  "no chs",
			input: [][]int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chans []<-chan int

			for _, slice := range tt.input {
				ch := make(chan int, len(slice))
				for _, v := range slice {
					ch <- v
				}
				close(ch)
				chans = append(chans, ch)
			}

			out := Merge(chans...)
			got := collect(out)

			if !sameElements(got, tt.want) {
				t.Fatalf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
