package main

import "fmt"

func main() {
	// var で型を明示して宣言
	var name string = "Gopher"
	// 短縮宣言（型推論）
	age := 10
	height := 1.75
	isActive := true
	// 宣言のみ（ゼロ値が入る）
	var score int

	fmt.Println("名前:", name)
	fmt.Println("年齢:", age)
	fmt.Println("身長:", height)
	fmt.Println("有効:", isActive)
	fmt.Println("スコア:", score)
}
