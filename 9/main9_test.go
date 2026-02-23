package main

import (
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	tests := []struct {
		name  string
		input []uint8
		want  []float64
	}{
		{
			name:  "basic",
			input: []uint8{1, 2, 3},
			want:  []float64{1, 8, 27},
		},
		{
			name:  "empty",
			input: []uint8{},
			want:  []float64{},
		},
		{
			name:  "1 value",
			input: []uint8{5},
			want:  []float64{125},
		},
		{
			name:  "zero",
			input: []uint8{0},
			want:  []float64{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan uint8)
			out := make(chan float64)

			Pipeline(in, out)

			go func() {
				for _, v := range tt.input {
					in <- v
				}
				close(in)
			}()

			got := make([]float64, 0)
			for v := range out {
				got = append(got, v)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Pipeline(%v) = %v, want %v",
					tt.input, got, tt.want)
			}
		})
	}
}
