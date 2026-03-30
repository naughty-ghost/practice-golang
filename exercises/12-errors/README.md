# 12 - エラーハンドリング

## 問題

以下の要件を満たすプログラムを作成してください。

1. `Divide(x, y float64) (float64, error)` 関数を作成する
2. `y` が `0` の場合はエラーを返す
3. `main` で正常ケースとエラーケースの両方を呼び出し、結果を出力する

## 出力例

```
10 / 3 = 3.33
エラー: 0で割ることはできません
```

## 解説

### Go のエラーハンドリング

Go には `try-catch` がありません。代わりに、関数の戻り値として `error` を返すのが慣例です。

```go
result, err := SomeFunction()
if err != nil {
    // エラー処理
    fmt.Println(err)
    return
}
// 正常処理
```

### error インターフェース

`error` は以下のインターフェースです。

```go
type error interface {
    Error() string
}
```

`fmt.Errorf()` や `errors.New()` で簡単にエラーを作成できます。

```go
import "fmt"
err := fmt.Errorf("値が不正です: %d", value)
```
