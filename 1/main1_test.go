package main

import (
	"reflect"
	"testing"
)

func TestGetType(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"int", 42, "int"},
		{"string", "Go", "string"},
		{"bool", true, "bool"},
		{"float64", 3.14, "float64"},
		{"complex", complex64(1 + 2i), "complex64"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetType(tt.input)
			if got != tt.want {
				t.Fatalf("GetType(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name   string
		input  []any
		expect string
	}{
		{
			"int+str",
			[]any{42, "Go", true},
			"42Gotrue",
		},
		{
			"complex",
			[]any{1 + 2i},
			"(1+2i)",
		},
		{
			"empty",
			[]any{},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToString(tt.input...)
			if got != tt.expect {
				t.Fatalf("ToString() = %q, want %q", got, tt.expect)
			}
		})
	}
}

func TestToRunes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []rune
	}{
		{"eng", "Go", []rune{'G', 'o'}},
		{"rus", "Пр", []rune{'П', 'р'}},
		{"empty", "", []rune{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToRunes(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ToRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestHashProperties(t *testing.T) {
	tests := []struct {
		name string
		str  string
		salt string
	}{
		{"eng", "Hello", "salt"},
		{"rus", "Привет", "go"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := []rune(tt.str)
			h1 := Hash(r, tt.salt)
			h2 := Hash(r, tt.salt)

			if h1 != h2 {
				t.Fatalf("Hash() not determ")
			}

			if len(h1) != 64 {
				t.Fatalf("Hash() length = %d, want 64", len(h1))
			}

			h3 := Hash([]rune(tt.str+"x"), tt.salt)
			if h1 == h3 {
				t.Fatalf("Hash() input change not reflected")
			}

			h4 := Hash(r, tt.salt+"x")
			if h1 == h4 {
				t.Fatalf("Hash() salt change not reflected")
			}
		})
	}
}
