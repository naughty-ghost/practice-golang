package main

import "fmt"

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("0で割ることはできません")
	}
	return x / y, nil
}

func main() {
	// 正常ケース
	result, err := Divide(10, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// エラーケース
	result, err = Divide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}
}
