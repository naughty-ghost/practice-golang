package main

import "fmt"

// 値渡し: 関数内で変更しても呼び出し元には影響しない
func Double(n int) {
	n = n * 2
}

// ポインタ渡し: 呼び出し元の変数を直接変更できる
func DoublePtr(n *int) {
	*n = *n * 2
}

func main() {
	x := 10
	fmt.Println("初期値:", x)

	Double(x)
	fmt.Println("Double後:", x)

	DoublePtr(&x)
	fmt.Println("DoublePtr後:", x)
}
