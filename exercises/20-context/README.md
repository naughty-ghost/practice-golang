# 20 - context によるキャンセル

## 問題

以下の要件を満たすプログラムを作成してください。

1. `longTask(ctx context.Context, id int)` 関数を作成する
   - 500 ミリ秒ごとに `"タスク {id}: 処理中..."` を出力する
   - `ctx.Done()` を受信したら `"タスク {id}: キャンセルされました"` を出力して終了する
2. `main` で `context.WithTimeout` を使い、2 秒のタイムアウト付き context を作成する
3. 3 つの goroutine で `longTask` を起動する
4. タイムアウト後、すべてのタスクが停止することを確認する
5. `"全タスク終了"` を出力する

## 出力例

```
タスク 1: 処理中...
タスク 2: 処理中...
タスク 3: 処理中...
タスク 1: 処理中...
タスク 3: 処理中...
タスク 2: 処理中...
タスク 1: 処理中...
タスク 2: 処理中...
タスク 3: 処理中...
タスク 2: キャンセルされました
タスク 1: キャンセルされました
タスク 3: キャンセルされました
全タスク終了
```

※ 出力順序は実行ごとに異なります。

## ヒント

- `longTask` 内でループを回し、`select` で `ctx.Done()` と `time.After` を待ち受ける
- `sync.WaitGroup` を使って全 goroutine の終了を待つ

## 解説

### context パッケージとは

`context` は goroutine の**キャンセル**、**タイムアウト**、**値の伝搬**を管理するパッケージです。Go のサーバーサイドプログラミングでは、ほぼすべての関数の第 1 引数に `ctx context.Context` を渡す慣習があります。

### context の生成

```go
// ベースとなる空の context
ctx := context.Background()

// キャンセル可能な context
ctx, cancel := context.WithCancel(parentCtx)
defer cancel()

// タイムアウト付き context
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()

// デッドライン付き context
ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))
defer cancel()
```

### キャンセルの検知

goroutine 内では `ctx.Done()` チャネルを `select` で監視します:

```go
func doWork(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("キャンセル:", ctx.Err())
            return
        default:
            // 通常の処理
        }
    }
}
```

`ctx.Err()` はキャンセルの理由を返します:
- `context.Canceled` — `cancel()` が呼ばれた
- `context.DeadlineExceeded` — タイムアウトまたはデッドラインを超過

### context の親子関係

context は**ツリー構造**で、親がキャンセルされると子もすべてキャンセルされます:

```
Background
  └── WithTimeout (5s)
        ├── WithCancel (リクエストA)
        └── WithCancel (リクエストB)
```

親の 5 秒タイムアウトが発火すると、リクエスト A, B も自動的にキャンセルされます。

### なぜ context が重要か

HTTP サーバーでクライアントが接続を切った場合、処理を続けるのはリソースの無駄です。`context` を使えば、不要になった処理を素早く停止できます。
