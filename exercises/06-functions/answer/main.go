package main

import "fmt"

// 2つの整数の合計を返す
func Add(x, y int) int {
	return x + y
}

// 2つの文字列を入れ替えて返す
func Swap(x, y string) (string, string) {
	return y, x
}

// 関数を引数に取る
func Apply(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	fmt.Println("Add:", Add(10, 20))

	a, b := Swap("hello", "world")
	fmt.Println("Swap:", a, b)

	fmt.Println("Apply (加算):", Apply(10, 20, func(a, b int) int {
		return a + b
	}))
	fmt.Println("Apply (乗算):", Apply(10, 20, func(a, b int) int {
		return a * b
	}))
}
