package main

import "fmt"

func FirstStep(number ...int) chan int {
	c := make(chan int)
	go func() {
		for item := range number {
			c <- number[item]
		}
		close(c)
	}()
	return c
}

func SecondStep(fromFirst chan int) chan int {
	c := make(chan int)
	go func() {
		//
		for item := range fromFirst {
			c <- item * item
		}
		close(c)
	}()
	return c
}

func main() {
	firstChan := FirstStep(1, 2, 3, 4, 5, 6)
	secondChan := SecondStep(firstChan)
	for item := range secondChan {
		fmt.Println(item)
	}
}
