package main

import "fmt"

func Diff(s1, s2 []string) []string {
	set := make(map[string]struct{}, len(s2))
	for _, v := range s2 {
		set[v] = struct{}{}
	}

	res := make([]string, 0, len(s1))
	for _, v := range s1 {
		if _, ok := set[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(Diff(slice1, slice2))
}
