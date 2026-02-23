package main

import (
	"reflect"
	"testing"
)

func TestAddAndExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 67)
	m.Add("b", 52)

	tests := []struct {
		key string
		want bool
	}{
		{"a", true},
		{"b", true},
		{"c", false},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			got := m.Exists(tt.key)
			if got != tt.want {
				t.Fatalf("Exists(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 52)

	tests := []struct {
		key string
		wantVal int
		wantOk bool
	}{
		{"a", 52, true},
		{"b", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			gotVal, gotOk := m.Get(tt.key)
			if gotVal != tt.wantVal || gotOk != tt.wantOk {
				t.Fatalf("Get(%q) = (%v,%v), want (%v,%v)", tt.key, gotVal, gotOk, tt.wantVal, tt.wantOk)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 67)
	m.Add("b", 52)

	m.Remove("a")
	if m.Exists("a") {
		t.Fatal("Remove() id not delete a")
	}

	m.Remove("c") 
	if m.Exists("b") != true {
		t.Fatal("Remove() delete random key")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 67)
	m.Add("b", 52)

	copied := m.Copy()
	if !reflect.DeepEqual(copied, map[string]int{"a":67,"b":52}) {
		t.Fatalf("Copy() = %v, want %v", copied, map[string]int{"a":67,"b":52})
	}

	m.Add("c", 100)
	if _, ok := copied["c"]; ok {
		t.Fatal("Copy() copy = modified original")
	}
}
