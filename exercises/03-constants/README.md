# 03 - 定数と iota

## 問題

以下の要件を満たすプログラムを作成してください。

1. 円周率 `Pi` を定数として `3.14159` で宣言する
2. 曜日を `iota` を使って定数で定義する（Sunday=0 から始まる）
3. `Pi` の値と、各曜日の値を出力する

## 出力例

```
Pi: 3.14159
Sunday: 0
Monday: 1
Tuesday: 2
Wednesday: 3
Thursday: 4
Friday: 5
Saturday: 6
```

## 解説

### iota

`iota` は `const` ブロック内で使える特殊な値で、0 から始まり行ごとに 1 ずつ増加します。
列挙型のように使うことができます。

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
)
```
