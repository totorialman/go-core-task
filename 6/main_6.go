package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generator(num int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		for i:=0;i<num;i++{
			ch <- r.Int()
		}
	}()

	return ch
}

func main() {
	for n := range Generator(-3){
		fmt.Println(n)
	}
}
