# 16 - goroutine と channel

## 問題

以下の要件を満たすプログラムを作成してください。

1. `square(n int, ch chan int)` 関数を作り、`n` の 2 乗を `ch` に送信する
2. `main` で 1 から 5 までそれぞれ goroutine で `square` を実行する
3. channel から 5 つの結果を受信して出力する

## 出力例

```
結果: 1
結果: 4
結果: 9
結果: 16
結果: 25
```

※ goroutine の実行順序は保証されないため、順番が異なる場合があります。

## 解説

### goroutine

`go` キーワードを関数呼び出しの前に付けると、その関数は新しい goroutine（軽量スレッド）で並行実行されます。

```go
go myFunction()  // 新しい goroutine で実行
```

### channel

channel は goroutine 間でデータをやりとりするための仕組みです。

```go
ch := make(chan int)    // チャネルの作成
ch <- 42               // 送信（ブロックする）
value := <-ch          // 受信（ブロックする）
```

送信と受信はどちらもブロックするため、goroutine 間の同期にも使えます。
