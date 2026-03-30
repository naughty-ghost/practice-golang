package main

import "fmt"

func main() {
	// 1から10まで
	fmt.Println("--- 1から10 ---")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// スライスの反復
	fmt.Println("--- スライス ---")
	langs := []string{"Go", "Python", "Rust"}
	for i, v := range langs {
		fmt.Printf("%d: %s\n", i, v)
	}

	// 偶数のみ（continueを使用）
	fmt.Println("--- 偶数のみ ---")
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
