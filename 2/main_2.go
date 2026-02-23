package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(s []int) []int {
	res := make([]int, 0, len(s))

	for _, v := range s {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func addElements(s []int, val int) []int {
	res := make([]int, len(s), len(s)+1)
	copy(res, s)
	res = append(res, val)
	return res
}

func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}

func removeElement(s []int, idx int) []int {
	if idx < 0 || idx >= len(s) {
		return copySlice(s)
	}
	res := make([]int, 0, len(s)-1)
	res = append(res, s[:idx]...)
	res = append(res, s[idx+1:]...)
	return res
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = r.Intn(100)
	}

	fmt.Println("original:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("even:", evenSlice)

	addedSlice := addElements(originalSlice, 67)
	fmt.Println("add 67:", addedSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("copy:", copiedSlice)

	originalSlice[0] = -1
	fmt.Println("modified original:", originalSlice)
	fmt.Println("copy:", copiedSlice)

	removedSlice := removeElement(originalSlice, 2)
	fmt.Println("removed element 2:", removedSlice)

}
