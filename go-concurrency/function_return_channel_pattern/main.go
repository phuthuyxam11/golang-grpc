package main

import "fmt"

func PrintSomething(message string) chan string {
	schan := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			schan <- fmt.Sprintf("%s_%d", message, i)
		}
	}()
	return schan
}

func main() {
	tmpChan := PrintSomething("hello_channel")
	for i := 0; i < 10; i++ {
		fmt.Println(<-tmpChan)
	}
}
