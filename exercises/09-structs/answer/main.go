package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

// 値レシーバ: フィールドを変更しないので値レシーバでOK
func (d Dog) Bark() {
	fmt.Printf("ワン! 私は%s、%d歳です\n", d.Name, d.Age)
}

// ポインタレシーバ: Ageを変更するのでポインタレシーバが必要
func (d *Dog) Birthday() {
	d.Age++
	fmt.Println("誕生日を迎えました")
}

func main() {
	dog := &Dog{Name: "ポチ", Age: 3}
	dog.Bark()
	dog.Birthday()
	dog.Bark()
}
