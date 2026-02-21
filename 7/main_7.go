package main

import (
	"fmt"
	"sync"
)

func Merge[T any](chs ...<-chan T) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}

	wg.Add(len(chs))
	for _, ch := range chs {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func main() {
	a := make(chan int, 10)
	b := make(chan int, 10)

	go func() {
		defer close(a)
		defer close(b)
		for i := range 10 {
			a <- i
			b <- i * 10
		}
	}()

	for i := range Merge(a, b) {
		fmt.Println(i)
	}
}
