package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name string
		input []int
		want []int
	}{
		{"all", []int{1, 2, 3, 4, 5}, []int{2, 4}},
		{"all even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"all odd", []int{1, 3, 5}, []int{}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("sliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name string
		input []int
		val int
		want []int
	}{
		{"add", []int{1,2,3}, 4, []int{1,2,3,4}},
		{"add to empty", []int{}, 5, []int{5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addElements(tt.input, tt.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("addElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1,2,3}
	copied := copySlice(original)
	
	if !reflect.DeepEqual(copied, original) {
		t.Fatal("copySlice() copied != original")
	}

	original[0] = 67
	if copied[0] == original[0] {
		t.Fatal("copySlice() copied = modified original")
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		input []int
		idx int
		want []int
	}{
		{"remove 2", []int{1,2,3,4}, 2, []int{1,2,4}},
		{"remove first", []int{1,2,3}, 0, []int{2,3}},
		{"remove last", []int{1,2,3}, 2, []int{1,2}},
		{"negative idx", []int{1,2,3}, -1, []int{1,2,3}},
		{"to big idx", []int{1,2,3}, 3, []int{1,2,3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.input, tt.idx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
