package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// === 一括読み取り ===
	fmt.Println("=== 一括読み取り ===")
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))

	// === 1行ずつ読み取り ===
	fmt.Println("\n=== 1行ずつ読み取り ===")
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("%d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
