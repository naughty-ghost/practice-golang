package main

import "fmt"

func main() {
	// 配列
	arr := [3]int{10, 20, 30}
	fmt.Println("配列:", arr)

	// スライスにappendで追加
	var sl []int
	sl = append(sl, 1, 2, 3, 4, 5)
	fmt.Println("スライス:", sl)

	// 3番目の要素（インデックス2）を削除
	sl = append(sl[:2], sl[3:]...)
	fmt.Println("削除後:", sl)

	// 逆順に並び替え
	nums := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println("逆順:", nums)
}
