package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// goroutine A: 1秒後に送信
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "A: 完了"
	}()

	// goroutine B: 2秒後に送信
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "B: 完了"
	}()

	// 先に届いた方を受信
	select {
	case msg := <-ch1:
		fmt.Println("受信:", msg)
	case msg := <-ch2:
		fmt.Println("受信:", msg)
	}

	// 残りの方を受信
	select {
	case msg := <-ch1:
		fmt.Println("受信:", msg)
	case msg := <-ch2:
		fmt.Println("受信:", msg)
	}
}
