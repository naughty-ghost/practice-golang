package main

import "fmt"

func main() {
	// マップの初期化
	stock := map[string]int{
		"apple":  3,
		"banana": 5,
		"orange": 2,
	}

	// すべてのキーと値を出力
	fmt.Println("--- 在庫一覧 ---")
	for k, v := range stock {
		fmt.Printf("%s: %d\n", k, v)
	}

	// 追加
	stock["grape"] = 8
	fmt.Println("grape を追加しました")

	// 削除
	delete(stock, "banana")
	fmt.Println("banana を削除しました")

	// 存在チェック
	if _, ok := stock["melon"]; ok {
		fmt.Println("melon は存在します")
	} else {
		fmt.Println("melon は存在しません")
	}

	// 最終結果
	fmt.Println("--- 最終在庫 ---")
	for k, v := range stock {
		fmt.Printf("%s: %d\n", k, v)
	}
}
