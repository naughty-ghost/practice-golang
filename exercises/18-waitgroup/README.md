# 18 - sync.WaitGroup

## 問題

以下の要件を満たすプログラムを作成してください。

1. `worker(id int, wg *sync.WaitGroup)` 関数を作成する
   - 関数内で `defer wg.Done()` を呼ぶ
   - `"Worker {id}: 開始"` を出力する
   - `time.Sleep` で 500 ミリ秒待機する
   - `"Worker {id}: 完了"` を出力する
2. `main` で `sync.WaitGroup` を使い、5 つの worker を goroutine で起動する
3. すべての worker が完了してから `"全ワーカー完了"` を出力する

## 出力例

```
Worker 1: 開始
Worker 2: 開始
Worker 3: 開始
Worker 4: 開始
Worker 5: 開始
Worker 3: 完了
Worker 1: 完了
Worker 5: 完了
Worker 2: 完了
Worker 4: 完了
全ワーカー完了
```

※ goroutine の実行順序は保証されないため、順番は異なる場合があります。

## ヒント

- `wg.Add(1)` は goroutine を起動する**前に**呼ぶ
- `wg` は**ポインタ**で渡す（値渡しするとコピーされて正しく動作しない）

## 解説

### sync.WaitGroup とは

`sync.WaitGroup` は、複数の goroutine の完了を待つための同期プリミティブです。channel で同じことはできますが、「結果は不要で、ただ完了を待ちたい」場合は WaitGroup の方がシンプルです。

```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Add(1)           // カウンタを +1
    go func(id int) {
        defer wg.Done() // カウンタを -1
        // 処理
    }(i)
}

wg.Wait() // カウンタが 0 になるまでブロック
```

### 3 つのメソッド

| メソッド | 説明 |
|---------|------|
| `Add(n)` | カウンタに `n` を加算する。goroutine 起動前に呼ぶ |
| `Done()` | カウンタを 1 減らす。`Add(-1)` と等価 |
| `Wait()` | カウンタが 0 になるまでブロックする |

### よくある間違い

```go
// NG: goroutine 内で Add を呼ぶと、Wait が先に通過する可能性がある
go func() {
    wg.Add(1)  // ← ここでは遅い
    defer wg.Done()
}()
wg.Wait()

// OK: goroutine 起動前に Add を呼ぶ
wg.Add(1)
go func() {
    defer wg.Done()
}()
wg.Wait()
```

### channel との使い分け

- **結果を受け取る必要がある** → channel
- **完了を待つだけでよい** → `sync.WaitGroup`
