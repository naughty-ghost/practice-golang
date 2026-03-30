# 15 - 文字列操作（正規表現とマルチバイト処理）

## 問題 1: 正規表現を使った置換

以下の文字列に対して、正規表現で置換を行ってください。

```
入力: "注文番号: A-1234, B-5678, C-9012 をご確認ください"
```

1. 注文番号のパターン（英大文字1文字 + ハイフン + 数字4桁）を `"***"` に置換して出力する
2. 同じパターンにマッチした注文番号を**すべて抽出**してスライスとして出力する

### 出力例

```
置換後: 注文番号: ***, ***, *** をご確認ください
抽出結果: [A-1234 B-5678 C-9012]
```

---

## 問題 2: マルチバイト文字列の指定文字数カット

以下の関数を作成してください。

```go
func truncate(s string, maxChars int) string
```

- 文字列 `s` を**先頭から `maxChars` 文字**に切り詰める
- `maxChars` が文字数以上の場合はそのまま返す
- マルチバイト文字（日本語など）を途中で壊さないこと

以下の入力でテストしてください:

```go
fmt.Println(truncate("Hello, 世界！れはテストです", 10))
fmt.Println(truncate("abcdefg", 5))
fmt.Println(truncate("短い", 10))
```

### 出力例

```
Hello, 世界！
abcde
短い
```

## ヒント

- 正規表現: `regexp.MustCompile` でコンパイル、`ReplaceAllString` で置換、`FindAllString` で抽出
- マルチバイト: `[]rune` に変換すると 1 文字 = 1 要素になる

## 解説

### 正規表現（regexp パッケージ）

Go の正規表現は `regexp` パッケージで提供されます。RE2 構文を採用しており、**常に線形時間**で実行されるのが特徴です（バックトラッキングしない）。

#### コンパイル

```go
// パターンが不正なら panic する（定数パターン向き）
re := regexp.MustCompile(`[A-Z]-\d{4}`)

// パターンが不正なら error を返す（動的パターン向き）
re, err := regexp.Compile(pattern)
```

#### よく使うメソッド

| メソッド | 説明 |
|---------|------|
| `MatchString(s)` | マッチするか判定（bool） |
| `FindString(s)` | 最初のマッチを返す |
| `FindAllString(s, n)` | 最大 n 個のマッチを返す（-1 で全件） |
| `ReplaceAllString(s, repl)` | マッチ部分を `repl` に置換 |
| `FindStringSubmatch(s)` | キャプチャグループ付きで最初のマッチを返す |

#### キャプチャグループの例

```go
re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
match := re.FindStringSubmatch("user@example.com")
// match[0] = "user@example.com"（全体マッチ）
// match[1] = "user"
// match[2] = "example"
// match[3] = "com"
```

#### 置換で後方参照を使う

```go
re := regexp.MustCompile(`([A-Z])-(\d{4})`)
result := re.ReplaceAllString("A-1234", "[$1:$2]")
// result = "[A:1234]"
```

---

### マルチバイト文字列の扱い

Go の `string` は **UTF-8 エンコードされたバイト列**です。日本語 1 文字は 3 バイトを占めるため、バイト単位の操作では文字が壊れます。

#### string, []byte, []rune の違い

```go
s := "Go言語"

len(s)          // 8  （バイト数: G=1, o=1, 言=3, 語=3）
len([]rune(s))  // 4  （文字数: G, o, 言, 語）
len([]byte(s))  // 8  （バイト数）
```

| 型 | 単位 | 用途 |
|---|---|---|
| `string` | バイト列 | 文字列の保持・比較 |
| `[]byte` | バイト | I/O、バイナリ処理 |
| `[]rune` | Unicode コードポイント | 文字単位の操作 |

#### 安全な文字数カット

```go
func truncate(s string, maxChars int) string {
    runes := []rune(s)
    if len(runes) <= maxChars {
        return s
    }
    return string(runes[:maxChars])
}
```

`[]rune` に変換することで、マルチバイト文字を壊さずにスライスできます。

#### for range による文字単位のイテレーション

`for range` は自動的に UTF-8 をデコードして `rune` 単位でイテレーションします:

```go
for i, r := range "Go言語" {
    fmt.Printf("バイト位置=%d, 文字=%c\n", i, r)
}
// バイト位置=0, 文字=G
// バイト位置=1, 文字=o
// バイト位置=2, 文字=言
// バイト位置=5, 文字=語
```

`i` はバイト位置なので 2 → 5 にジャンプしていることに注目してください（「言」が 3 バイトを占めるため）。

#### unicode/utf8 パッケージ

`[]rune` への変換なしで文字数を取得できます:

```go
utf8.RuneCountInString("Go言語") // 4
```
