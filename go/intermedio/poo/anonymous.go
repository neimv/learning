package main

import (
	"fmt"
	"time"
)

func sum(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)

	c := make(chan int)
	go func() {
		fmt.Println("Starting function")
		time.Sleep(1 * time.Second)
		fmt.Println("End")
		c <- 1
	}()

	<-c

	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4))

	printNames("Alice", "Bob", "Charlie", "Dave")

	fmt.Println(getValues(2))
}
