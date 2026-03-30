package main

import "fmt"

func main() {
	score := 85

	// if版
	var grade string
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}
	fmt.Println("if版:", grade)

	// switch版
	var grade2 string
	switch {
	case score >= 90:
		grade2 = "A"
	case score >= 80:
		grade2 = "B"
	case score >= 70:
		grade2 = "C"
	default:
		grade2 = "D"
	}
	fmt.Println("switch版:", grade2)
}
