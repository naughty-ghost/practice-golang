package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// === 新規作成して書き込み ===
	content := "1行目: Hello, Go!\n2行目: ファイル書き込みのテスト\n3行目: 完了\n"
	err := os.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 書き込み結果を確認
	fmt.Println("=== 書き込み後 ===")
	data, err := os.ReadFile("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))

	// === 追記 ===
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "4行目: 追記されました")
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	// 追記結果を確認
	fmt.Println("\n=== 追記後 ===")
	data, err = os.ReadFile("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
