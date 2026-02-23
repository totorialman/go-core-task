package main

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	tests := []struct {
		name      string
		num       int
		wantCount int
	}{
		{
			name:      "generate 5 numbers",
			num:       5,
			wantCount: 5,
		},
		{
			name:      "generate 0 numbers",
			num:       0,
			wantCount: 0,
		},
		{
			name:      "generate -3 numbers",
			num:       -3,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := Generator(tt.num)

			count := 0
			for range ch {
				count++
			}

			if count != tt.wantCount {
				t.Fatalf("Generator(%d) = %d, want %d",
					tt.num, count, tt.wantCount)
			}
		})
	}
}
