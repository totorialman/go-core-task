package main

import (
	"fmt"
)

func Pipeline(inCh <-chan uint8, outCh chan<- float64) {
	go func() {
		defer close(outCh)
		for n := range inCh {
			outCh <- float64(n) * float64(n) * float64(n)
		}
	}()
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	Pipeline(in, out)

	go func() {
		defer close(in)
		for i := uint8(0); i < 10; i++ {
			in <- i
		}
	}()

	for res := range out {
		fmt.Println(res)
	}
}
