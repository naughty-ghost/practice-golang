package main

import (
	"fmt"
	"syscall/js"
)

func add(this js.Value, args []js.Value) any {
	a := args[0].Int()
	b := args[1].Int()
	return a + b
}

func main() {
	fmt.Println("Hello from Go WebAssembly!")

	// JavaScript から呼べる関数を登録
	js.Global().Set("goAdd", js.FuncOf(add))

	// プログラムを終了させない
	select {}
}
