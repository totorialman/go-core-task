package main

import "fmt"

func Intersection(a, b []int) (bool, []int) {
	if len(a) > len(b) {
		a, b = b, a
	}

	set := make(map[int]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	var res []int
	seen := make(map[int]struct{})

	for _, v := range b {
		if _, ok := set[v]; ok {
			if _, added := seen[v]; !added {
				res = append(res, v)
				seen[v] = struct{}{}
			}
		}
	}

	if len(res) == 0{
		return false, []int{}
	}

	return true, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(Intersection(a, b))
}
