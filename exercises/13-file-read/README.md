# 13 - ファイル読み取り

## 問題

以下の要件を満たすプログラムを作成してください。

### 準備

同じディレクトリに `sample.txt` を作成し、以下の内容を書いてください:

```
Go言語は2009年にGoogleで開発されました。
シンプルな構文と高い並行処理性能が特徴です。
静的型付け言語でありながら、動的言語のような書きやすさを持ちます。
```

### 要件

1. `sample.txt` を**一括読み取り**して内容を全て出力する（`os.ReadFile` を使用）
2. `sample.txt` を**1行ずつ読み取り**して、行番号付きで出力する（`bufio.Scanner` を使用）

## 出力例

```
=== 一括読み取り ===
Go言語は2009年にGoogleで開発されました。
シンプルな構文と高い並行処理性能が特徴です。
静的型付け言語でありながら、動的言語のような書きやすさを持ちます。

=== 1行ずつ読み取り ===
1: Go言語は2009年にGoogleで開発されました。
2: シンプルな構文と高い並行処理性能が特徴です。
3: 静的型付け言語でありながら、動的言語のような書きやすさを持ちます。
```

## ヒント

- `os.ReadFile` はファイル全体を `[]byte` で返す
- `bufio.NewScanner` + `scanner.Scan()` で 1 行ずつ読める
- ファイルを `os.Open` で開いたら `defer f.Close()` で閉じる

## 解説

### ファイル読み取りの 2 つの方法

Go ではファイルの読み取りに主に 2 つのアプローチがあります。

#### 1. 一括読み取り（os.ReadFile）

ファイル全体をメモリに読み込みます。小〜中サイズのファイルに向いています。

```go
data, err := os.ReadFile("sample.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

`os.ReadFile` はファイルのオープン・読み取り・クローズを全て行うため、最もシンプルな方法です。

#### 2. 行単位の読み取り（bufio.Scanner）

大きなファイルや、1 行ずつ処理したい場合に使います。

```go
f, err := os.Open("sample.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

### os.Open vs os.OpenFile

| 関数 | 用途 |
|------|------|
| `os.Open(name)` | 読み取り専用で開く（`O_RDONLY`） |
| `os.OpenFile(name, flag, perm)` | フラグとパーミッションを指定して開く |

`os.Open` は `os.OpenFile(name, os.O_RDONLY, 0)` のショートカットです。

### defer によるリソース解放

ファイルを開いたら必ず閉じる必要があります。`defer` を使うと、関数の終了時に自動的に `Close()` が呼ばれます。

```go
f, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close() // 関数終了時に必ず実行される
```

`defer` はエラーで早期リターンした場合でも確実に実行されるため、リソースリークを防げます。

### エラーハンドリングの注意点

ファイル操作は必ずエラーチェックを行います。ファイルが存在しない、権限がないなどの問題が実行時に発生するためです。

```go
data, err := os.ReadFile("notfound.txt")
if err != nil {
    // os.IsNotExist(err) でファイルの存在チェックも可能
    if os.IsNotExist(err) {
        fmt.Println("ファイルが見つかりません")
    }
    log.Fatal(err)
}
```
