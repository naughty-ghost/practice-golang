# 06 - 関数

## 問題

以下の関数を作成し、`main` から呼び出して結果を出力してください。

1. `Add(x, y int) int` : 2 つの整数の合計を返す
2. `Swap(x, y string) (string, string)` : 2 つの文字列を入れ替えて返す（複数戻り値）
3. `Apply(x, y int, op func(int, int) int) int` : 2 つの整数に関数 `op` を適用した結果を返す

## 出力例

```
Add: 30
Swap: world hello
Apply (加算): 30
Apply (乗算): 200
```

## 解説

### 複数戻り値

Go の関数は複数の値を返すことができます。エラーハンドリングで頻繁に使われるパターンです。

```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("0で割ることはできません")
    }
    return x / y, nil
}
```

### 関数は第一級オブジェクト

Go では関数を変数に代入したり、引数として渡すことができます。
