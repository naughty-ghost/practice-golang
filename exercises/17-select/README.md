# 17 - select 文

## 問題

以下の要件を満たすプログラムを作成してください。

1. 2 つの channel `ch1`（string 型）と `ch2`（string 型）を作成する
2. goroutine A: 1 秒後に `ch1` へ `"A: 完了"` を送信する
3. goroutine B: 2 秒後に `ch2` へ `"B: 完了"` を送信する
4. `select` を使って、**先に届いた方**のメッセージを出力する
5. さらにもう一度 `select` で、**残りの方**のメッセージも出力する

## 出力例

```
受信: A: 完了
受信: B: 完了
```

## ヒント

- `time.Sleep` で送信を遅延させる
- `select` は複数の `case` のうち、準備ができたものを実行する

## 解説

### select 文とは

`select` は複数の channel 操作を同時に待ち受けるための構文です。`switch` に似ていますが、対象が channel の送受信である点が異なります。

```go
select {
case msg := <-ch1:
    fmt.Println("ch1 から受信:", msg)
case msg := <-ch2:
    fmt.Println("ch2 から受信:", msg)
}
```

- 複数の `case` が同時に準備完了の場合、**ランダムに 1 つが選ばれる**
- どの `case` も準備できていない場合、**ブロック**して待つ
- `default` を書くとブロックせずに即座に `default` が実行される（ノンブロッキング受信）

### ノンブロッキング受信

```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("まだデータがない")
}
```

### タイムアウトパターン

`select` と `time.After` を組み合わせると、タイムアウトを簡単に実装できます。

```go
select {
case result := <-ch:
    fmt.Println("結果:", result)
case <-time.After(3 * time.Second):
    fmt.Println("タイムアウト")
}
```

この仕組みにより、goroutine が永遠に待ち続けることを防げます。
