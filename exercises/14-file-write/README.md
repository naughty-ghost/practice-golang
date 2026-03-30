# 14 - ファイル書き込み

## 問題

以下の要件を満たすプログラムを作成してください。

1. `output.txt` に以下の 3 行を**新規作成して書き込む**（`os.WriteFile` を使用）

```
1行目: Hello, Go!
2行目: ファイル書き込みのテスト
3行目: 完了
```

2. 書き込んだ `output.txt` の内容を読み取って出力し、正しく書き込まれたことを確認する

3. `output.txt` に `"4行目: 追記されました"` を**追記する**（`os.OpenFile` で `os.O_APPEND` を使用）

4. 再度 `output.txt` を読み取って、追記後の内容を出力する

## 出力例

```
=== 書き込み後 ===
1行目: Hello, Go!
2行目: ファイル書き込みのテスト
3行目: 完了

=== 追記後 ===
1行目: Hello, Go!
2行目: ファイル書き込みのテスト
3行目: 完了
4行目: 追記されました
```

## ヒント

- `os.WriteFile` で一括書き込み（ファイルが既存なら上書き）
- `os.OpenFile` で追記モードで開くには `os.O_APPEND|os.O_WRONLY` を指定する
- パーミッションは `0644` が一般的

## 解説

### ファイル書き込みの方法

#### 1. 一括書き込み（os.WriteFile）

ファイル全体を一度に書き込みます。ファイルが存在すれば上書き、なければ新規作成されます。

```go
content := []byte("Hello, Go!\n")
err := os.WriteFile("output.txt", content, 0644)
if err != nil {
    log.Fatal(err)
}
```

第 3 引数の `0644` はファイルパーミッション（所有者が読み書き可、その他は読み取りのみ）です。

#### 2. 追記（os.OpenFile + O_APPEND）

既存ファイルの末尾にデータを追加します。

```go
f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

_, err = fmt.Fprintln(f, "追記する行")
if err != nil {
    log.Fatal(err)
}
```

#### 3. bufio.Writer によるバッファ付き書き込み

大量のデータを書き込む場合、`bufio.Writer` でバッファリングすると効率的です。

```go
f, err := os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

w := bufio.NewWriter(f)
fmt.Fprintln(w, "バッファ付き書き込み")
w.Flush() // バッファの内容をファイルに書き出す
```

`Flush()` を忘れるとデータがファイルに書き込まれないので注意してください。

### OpenFile のフラグ

`os.OpenFile` の第 2 引数にはフラグをビット OR で組み合わせて指定します。

| フラグ | 説明 |
|-------|------|
| `os.O_RDONLY` | 読み取り専用 |
| `os.O_WRONLY` | 書き込み専用 |
| `os.O_RDWR` | 読み書き両用 |
| `os.O_CREATE` | ファイルが存在しなければ新規作成 |
| `os.O_APPEND` | 書き込み時に末尾に追加 |
| `os.O_TRUNC` | 既存ファイルの内容をクリアして開く |

よく使う組み合わせ:

```go
// 新規作成 or 上書き
os.OpenFile("f.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

// 新規作成 or 追記
os.OpenFile("f.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
```

### パーミッション

Unix 系のファイルパーミッションを 8 進数で指定します。

| 値 | 意味 |
|----|------|
| `0644` | 所有者: 読み書き / グループ・その他: 読み取りのみ |
| `0755` | 所有者: 全権限 / グループ・その他: 読み取り+実行 |
| `0600` | 所有者のみ読み書き（機密ファイル向き） |

### fmt.Fprint 系関数

`fmt.Println` などの出力先をファイルに切り替えられます。

```go
fmt.Fprintln(f, "ファイルに出力")   // 改行付き
fmt.Fprintf(f, "値: %d\n", 42)     // フォーマット付き
fmt.Fprint(f, "改行なし")           // 改行なし
```

第 1 引数に `io.Writer` インターフェースを満たす値を渡します。`os.File` は `io.Writer` を実装しているため、そのまま渡せます。
