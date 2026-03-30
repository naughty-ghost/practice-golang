# 22 - WebAssembly

## 問題

Go のコードを WebAssembly（Wasm）にコンパイルし、ブラウザ上で動かしてください。

### 要件

1. `main.go` に以下の機能を実装する
   - ページ読み込み時にコンソールへ `"Hello from Go WebAssembly!"` を出力する
   - JavaScript 側から呼び出せる `add(a, b)` 関数を Go で登録し、結果を返す
   - HTML のボタンをクリックすると `add(3, 5)` の結果がページに表示される
2. `GOOS=js GOARCH=wasm` でビルドする
3. `index.html` を作成してブラウザで動作確認する

### ビルドと実行

```bash
# Wasm にコンパイル
GOOS=js GOARCH=wasm go build -o main.wasm main.go

# Go に同梱されている wasm_exec.js をコピー
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .

# ローカルサーバーで配信（Python 3 の場合）
python3 -m http.server 8080
```

ブラウザで `http://localhost:8080` を開いてください。

## 出力例

ブラウザのコンソール:
```
Hello from Go WebAssembly!
```

ボタンクリック後、ページ上に:
```
結果: 8
```

## ヒント

- `syscall/js` パッケージで JavaScript と連携する
- `js.Global()` でグローバルオブジェクト（`window`）にアクセスできる
- `js.FuncOf` で Go の関数を JavaScript から呼べるようにする
- `main` 関数を終了させないために `select {}` で永久ブロックする

## 解説

### WebAssembly（Wasm）とは

WebAssembly はブラウザ上で高速に実行できるバイナリフォーマットです。Go は標準で Wasm へのクロスコンパイルをサポートしており、追加ツールなしで Wasm バイナリを生成できます。

### ビルド方法

環境変数 `GOOS=js GOARCH=wasm` を指定するだけです。

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

ブラウザで Wasm を読み込むには、Go が提供する `wasm_exec.js`（グルーコード）が必要です。

```bash
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" .
```

### syscall/js パッケージ

Go から JavaScript の世界にアクセスするためのパッケージです。

#### js.Global() — グローバルオブジェクト

```go
// console.log を呼ぶ
js.Global().Get("console").Call("log", "Hello from Go!")

// document.getElementById
doc := js.Global().Get("document")
elem := doc.Call("getElementById", "result")
elem.Set("innerText", "Go から書き換え")
```

#### js.FuncOf — Go 関数を JS に公開

```go
addFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
    a := args[0].Int()
    b := args[1].Int()
    return a + b
})
js.Global().Set("goAdd", addFunc)
```

これにより、JavaScript 側から `goAdd(3, 5)` のように呼び出せます。

`js.FuncOf` のシグネチャ:
- `this` — JavaScript の `this` 値
- `args` — JavaScript から渡された引数（`js.Value` のスライス）
- 戻り値 — `any` 型（`int`, `string`, `bool`, `nil`, `js.Value` など）

#### js.Value の型変換

| メソッド | 説明 |
|---------|------|
| `Int()` | int に変換 |
| `Float()` | float64 に変換 |
| `String()` | string に変換 |
| `Bool()` | bool に変換 |
| `Get(name)` | プロパティを取得 |
| `Set(name, value)` | プロパティを設定 |
| `Call(name, args...)` | メソッドを呼び出し |

### main 関数を終了させない

Wasm の Go プログラムは `main` が終了するとすべての登録した関数も無効になります。`select {}` で永久にブロックさせて、プログラムを生かし続けます。

```go
func main() {
    // 関数の登録など
    js.Global().Set("goAdd", js.FuncOf(addFunc))

    // プログラムを終了させない
    select {}
}
```

### HTML 側の構成

```html
<script src="wasm_exec.js"></script>
<script>
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
  });
</script>
```

1. `wasm_exec.js` を読み込む（Go ランタイムのグルーコード）
2. `Go` オブジェクトを生成
3. `WebAssembly.instantiateStreaming` で `.wasm` ファイルを読み込み・コンパイル
4. `go.run()` で Go の `main` 関数を実行

### 注意点

- Wasm ファイルの配信には `Content-Type: application/wasm` が必要。`python3 -m http.server` は自動で対応する
- Wasm バイナリは Go ランタイムを含むためサイズが大きくなる（数 MB）。サイズが問題になる場合は [TinyGo](https://tinygo.org/) を検討する
- `syscall/js` は Wasm ビルド専用のため、通常の `go run` では動作しない
