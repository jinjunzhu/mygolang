package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	close(ch1)
	for elem1 := range ch1 {
		fmt.Printf("The first element received from channel ch1: %v\n", elem1)
	}

}
