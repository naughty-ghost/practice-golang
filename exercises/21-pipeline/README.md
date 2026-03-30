# 21 - パイプラインパターン

## 問題

以下の要件を満たす**パイプライン**を構築してください。

1. `generate(nums ...int) <-chan int` — 引数の数値を channel に送信する（送信完了後 close する）
2. `square(in <-chan int) <-chan int` — 受信した値を 2 乗して新しい channel に送信する
3. `filter(in <-chan int, threshold int) <-chan int` — `threshold` 以上の値のみ新しい channel に送信する
4. `main` でこれらを接続し、1〜10 を入力して、2 乗した結果のうち 20 以上のものだけを出力する

## 出力例

```
25
36
49
64
81
100
```

（5^2=25, 6^2=36, ..., 10^2=100 が 20 以上）

## ヒント

- 各ステージは goroutine で動作し、channel でつながる
- 送信側が `close(ch)` すると、受信側の `for range` ループが終了する

## 解説

### パイプラインパターンとは

パイプラインは、複数の**ステージ**を channel で直列に接続する並行処理パターンです。UNIX のパイプ（`cat file | grep error | wc -l`）と同じ考え方です。

```
generate → square → filter → 出力
```

各ステージは独立した goroutine で動作し、channel を介してデータを流します。

### ステージの書き方

各ステージは以下のパターンに従います:

```go
func stage(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out) // 入力が終わったら出力も閉じる
        for v := range in {
            out <- process(v) // 変換して送信
        }
    }()
    return out
}
```

ポイント:
- **入力**: `<-chan T`（受信専用チャネル）
- **出力**: `<-chan T`（受信専用チャネルとして返す）
- **close**: 処理が終わったら `close(out)` する。これにより下流の `for range` が終了する

### channel の方向指定

```go
chan int     // 送受信可能
<-chan int   // 受信専用（読み取りのみ）
chan<- int   // 送信専用（書き込みのみ）
```

方向を制限することで、誤った使い方をコンパイル時に防げます。

### for range によるチャネル受信

```go
for v := range ch {
    // ch が close されるまでループ
    fmt.Println(v)
}
```

`close(ch)` されると `for range` は自動的に終了します。close しないとデッドロックになるので注意してください。

### パイプラインの接続

```go
func main() {
    ch := generate(1, 2, 3, 4, 5)
    sq := square(ch)
    result := filter(sq, 10)

    for v := range result {
        fmt.Println(v)
    }
}
```

各ステージが goroutine で並行動作するため、データが流れ始めると全ステージが同時に処理を行います。これにより、大量データの処理でもメモリ効率が良くなります。
