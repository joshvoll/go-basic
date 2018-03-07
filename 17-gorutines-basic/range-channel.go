package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// for func
	go func() {
		for i := 0; i < 45; i++ {
			c <- i
		}
		close(c)
	}()

	// looping the channle
	for m := range c {
		fmt.Println(m)
	}
}