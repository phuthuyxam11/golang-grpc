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

func FanIn(chan1, chan2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case <-chan1:
				c <- <-chan1
			case <-chan2:
				c <- <-chan2
			}
		}
	}()
	return c
}

func main() {
	/*
		nếu có nhiều goroutine được khai báo trong 1 hàm main sẽ sảy ra hiện tượng channel blocking
		=> luôn ưu tiên chạy cho goroutine được khai báo trước. phải đợi hết routine 1 đc gửi và nhận data, thì các
		routine khác mới đc chạy => fanin pattern. -> channel nào có data trước thì chạy trước
	*/
	coffee := PrintSomething("coffee")
	bread := PrintSomething("bread")

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(<-coffee)
			fmt.Println(<-bread)
		}
		result:
		coffee_0
		bread_0
		coffee_1
		bread_1
		coffee_2
		bread_2
		coffee_3
		bread_3
		...
		coffee_9
		bread_9

	*/
	serve := FanIn(coffee, bread)
	for i := 0; i < 10; i++ {
		fmt.Println(<-serve)
	}
	fmt.Println("main finish.")
}
