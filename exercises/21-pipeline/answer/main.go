package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func filter(in <-chan int, threshold int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v >= threshold {
				out <- v
			}
		}
	}()
	return out
}

func main() {
	ch := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sq := square(ch)
	result := filter(sq, 20)

	for v := range result {
		fmt.Println(v)
	}
}
