# 19 - sync.Mutex

## 問題

以下の要件を満たすプログラムを作成してください。

1. `Counter` 構造体を作成する（フィールド: `mu sync.Mutex`, `value int`）
2. `(c *Counter) Increment()` メソッドを作成し、Mutex でロックしてから `value` を +1 する
3. `(c *Counter) Value() int` メソッドを作成し、Mutex でロックしてから `value` を返す
4. `main` で 1000 個の goroutine を起動し、それぞれが `Increment()` を 1 回呼ぶ
5. すべての goroutine 完了後に `Counter` の値を出力する（常に `1000` になるべき）

## 出力例

```
カウンタの値: 1000
```

## ヒント

- Mutex なしだと競合（race condition）が起き、結果が 1000 にならない
- `go run -race main.go` でデータ競合を検出できる

## 解説

### データ競合（Race Condition）とは

複数の goroutine が同時に同じ変数を読み書きすると、結果が不定になります。

```go
// 危険な例: 複数の goroutine が同時に counter++ を実行
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        counter++ // 読み取り → 加算 → 書き込み の3ステップが割り込まれる
    }()
}
// counter は 1000 にならないことがある
```

`counter++` は内部的に「読み取り → 加算 → 書き込み」の 3 ステップで、途中に別の goroutine が割り込む可能性があります。

### sync.Mutex

`Mutex`（Mutual Exclusion = 相互排他）は、クリティカルセクション（同時にアクセスしてはいけない区間）を保護します。

```go
var mu sync.Mutex
var counter int

mu.Lock()       // ロックを取得（他の goroutine はここで待つ）
counter++       // 安全に操作
mu.Unlock()     // ロックを解放
```

`defer` と組み合わせるのが一般的です:

```go
func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

### sync.RWMutex

読み取りが多く書き込みが少ない場合、`RWMutex` を使うとパフォーマンスが向上します。

```go
var mu sync.RWMutex

mu.RLock()    // 読み取りロック（複数の goroutine が同時に取得可能）
mu.RUnlock()

mu.Lock()     // 書き込みロック（排他的。読み取りロックとも排他）
mu.Unlock()
```

### Race Detector

Go には組み込みの競合検出ツールがあります:

```bash
go run -race main.go
go test -race ./...
```

Mutex を外して `-race` 付きで実行すると、データ競合が報告されるので試してみてください。
