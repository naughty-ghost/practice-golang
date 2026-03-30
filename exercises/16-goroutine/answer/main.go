package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int)

	for i := 1; i <= 5; i++ {
		go square(i, ch)
	}

	// 5つの結果を受信
	for i := 0; i < 5; i++ {
		result := <-ch
		fmt.Println("結果:", result)
	}
}
