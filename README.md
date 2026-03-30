# Go 基本構文チュートリアル

Go言語の基本構文を問題形式で学ぶチュートリアルリポジトリです。

## 使い方

1. `exercises/` 内の各フォルダにある `README.md` を読み、問題に取り組む
2. `main.go` を自分で書いてみる
3. `go run main.go` で実行し、出力例と比較する
4. わからない場合は `answer/main.go` の模範解答を参考にする

## 問題一覧

| #   | テーマ                                          | 内容                                 |
| --- | ----------------------------------------------- | ------------------------------------ |
| 01  | [Hello World](exercises/01-hello-world/)        | プログラムの基本構造、`fmt.Println`  |
| 02  | [変数と型](exercises/02-variables/)             | 変数宣言、基本型、ゼロ値             |
| 03  | [定数と iota](exercises/03-constants/)          | `const`、`iota` による列挙           |
| 04  | [if / switch](exercises/04-if-switch/)          | 条件分岐、簡易文付き if              |
| 05  | [for ループ](exercises/05-for-loop/)            | `for`、`range`、`break` / `continue` |
| 06  | [関数](exercises/06-functions/)                 | 関数定義、複数戻り値、無名関数       |
| 07  | [配列とスライス](exercises/07-slices/)          | 配列、スライス操作、`append`         |
| 08  | [マップ](exercises/08-maps/)                    | `map` の基本操作、存在チェック       |
| 09  | [構造体](exercises/09-structs/)                 | 構造体の定義とメソッド               |
| 10  | [インターフェース](exercises/10-interfaces/)    | インターフェースの定義と実装         |
| 11  | [ポインタ](exercises/11-pointers/)              | ポインタの基本、参照渡し             |
| 12  | [エラーハンドリング](exercises/12-errors/)      | `error` 型、独自エラー               |
| 13  | [ファイル読み取り](exercises/13-file-read/)     | `os.ReadFile`、`bufio.Scanner`       |
| 14  | [ファイル書き込み](exercises/14-file-write/)    | `os.WriteFile`、`os.OpenFile`、追記  |
| 15  | [文字列操作](exercises/15-string-operations/)   | 正規表現、マルチバイト文字のカット   |
| 16  | [goroutine と channel](exercises/16-goroutine/) | 並行処理の基本                       |
| 17  | [select 文](exercises/17-select/)               | 複数チャネルの待ち受け、タイムアウト |
| 18  | [sync.WaitGroup](exercises/18-waitgroup/)       | 複数 goroutine の完了待ち            |
| 19  | [sync.Mutex](exercises/19-mutex/)               | データ競合の防止、Race Detector      |
| 20  | [context](exercises/20-context/)                | キャンセル、タイムアウト、親子伝搬   |
| 21  | [パイプライン](exercises/21-pipeline/)          | ステージの直列接続、チャネル方向指定 |

## 実行環境

- Go 1.19 以上
- インストール: https://go.dev/doc/install

