package main

import (
	"fmt"
	"regexp"
)

// truncate はマルチバイト文字を壊さずに指定文字数で切り詰める
func truncate(s string, maxChars int) string {
	runes := []rune(s)
	if len(runes) <= maxChars {
		return s
	}
	return string(runes[:maxChars])
}

func main() {
	// === 問題 1: 正規表現を使った置換 ===
	input := "注文番号: A-1234, B-5678, C-9012 をご確認ください"

	re := regexp.MustCompile(`[A-Z]-\d{4}`)

	// 置換
	replaced := re.ReplaceAllString(input, "***")
	fmt.Println("置換後:", replaced)

	// 抽出
	matches := re.FindAllString(input, -1)
	fmt.Println("抽出結果:", matches)

	// === 問題 2: マルチバイト文字列の指定文字数カット ===
	fmt.Println(truncate("Hello, 世界！これはテストです", 10))
	fmt.Println(truncate("abcdefg", 5))
	fmt.Println(truncate("短い", 10))
}
