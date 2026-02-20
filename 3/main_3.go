package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copied := make(map[string]int, len(m.data))
	for key, val := range m.data {
		copied[key] = val
	}
	return copied
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.data[key]
	return val, ok
}

func main() {
	m := NewStringIntMap()

	m.Add("artur", 67)
	m.Add("друн", 52)
	fmt.Println("add:", m.data)

	fmt.Println("exists artur:", m.Exists("artur"))
	fmt.Println("exists друн:", m.Exists("друн"))

	val, ok := m.Get("друн")
	fmt.Println("get друн:", val, ok)

	m.Remove("artur")
	fmt.Println("remove artur:", m.data)

	copied := m.Copy()
	fmt.Println("copy:", copied)

	m.Add("ч", 30)
	fmt.Println("original:", m.data)
	fmt.Println("copy after original modified:", copied)
}
