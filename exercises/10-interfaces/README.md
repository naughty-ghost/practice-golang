# 10 - インターフェース

## 問題

以下の要件を満たすプログラムを作成してください。

1. `Shape` インターフェースを定義する（メソッド: `Area() float64`）
2. `Circle` 構造体（フィールド: `Radius float64`）に `Area()` を実装する
3. `Rectangle` 構造体（フィールド: `Width, Height float64`）に `Area()` を実装する
4. `PrintArea(s Shape)` 関数を作り、面積を出力する
5. `Circle` と `Rectangle` をそれぞれ `PrintArea` に渡して出力する

## 出力例

```
面積: 78.54
面積: 24.00
```

## 解説

### 暗黙的なインターフェース実装

Go では `implements` キーワードがありません。インターフェースが要求するメソッドをすべて実装していれば、自動的にそのインターフェースを満たします。

```go
type Writer interface {
    Write(data []byte) (int, error)
}

// File が Write メソッドを持っていれば Writer を満たす
type File struct{ /* ... */ }
func (f *File) Write(data []byte) (int, error) { /* ... */ }
```

### インターフェース設計のコツ

インターフェースのメソッドは少なく保つのが Go の慣習です。標準ライブラリの `io.Reader` や `io.Writer` はメソッドが 1 つだけです。
