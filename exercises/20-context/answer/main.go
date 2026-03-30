package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func longTask(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("タスク %d: キャンセルされました\n", id)
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("タスク %d: 処理中...\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go longTask(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("全タスク終了")
}
