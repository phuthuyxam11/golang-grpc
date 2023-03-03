package main

import (
	"fmt"
	"go.concurrency.example/example"
	"log"
)

func main() {
	result, err := example.FindCountText()
	if err != nil {
		log.Fatalln("error: ", err)
		return
	}
	fmt.Println(result)
}
