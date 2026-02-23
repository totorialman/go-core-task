package main

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		s1   []string
		s2   []string
		want []string
	}{
		{
			name: "difference",
			s1:   []string{"apple", "banana", "cherry"},
			s2:   []string{"banana"},
			want: []string{"apple", "cherry"},
		},
		{
			name: "no overlap",
			s1:   []string{"a", "b"},
			s2:   []string{"c", "d"},
			want: []string{"a", "b"},
		},
		{
			name: "overlap",
			s1:   []string{"a", "b"},
			s2:   []string{"a", "b"},
			want: []string{},
		},
		{
			name: "empty s1",
			s1:   []string{},
			s2:   []string{"a"},
			want: []string{},
		},
		{
			name: "empty s2",
			s1:   []string{"a", "b"},
			s2:   []string{},
			want: []string{"a", "b"},
		},
		{
			name: "empty all",
			s1:   []string{},
			s2:   []string{},
			want: []string{},
		},
		{
			name: "duplicates in s1",
			s1:   []string{"a", "b", "a"},
			s2:   []string{"b"},
			want: []string{"a", "a"},
		},
		{
			name: "duplicates in s2",
			s1:   []string{"a", "b", "c"},
			s2:   []string{"b", "b", "b"},
			want: []string{"a", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.s1, tt.s2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Diff(%v, %v) = %v, want %v",
					tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
