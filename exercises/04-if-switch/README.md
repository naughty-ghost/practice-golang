# 04 - if / switch

## 問題

整数の点数 `score` を `85` として宣言し、以下のルールで評価を出力してください。

- 90 以上: `"A"`
- 80 以上: `"B"`
- 70 以上: `"C"`
- 70 未満: `"D"`

**追加課題**: 同じ処理を `switch` 文でも書いてみてください。

## 出力例

```
if版: B
switch版: B
```

## 解説

### 簡易文付き if

Go の `if` は条件の前に簡易文を書けます。変数のスコープが `if` ブロック内に限定されるため便利です。

```go
if x := compute(); x > 0 {
    fmt.Println("正の値:", x)
}
// ここでは x は使えない
```

### switch の条件なし記法

`switch` に条件式を書かず、各 `case` に条件を書く方法があります。

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
}
```
