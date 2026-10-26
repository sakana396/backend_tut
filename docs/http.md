# HTTP サーバ演習

[研修1のトップに戻る](../README.md)


Go の標準ライブラリで小さな HTTP サーバを作ります。
すべてのコマンドは、`go.mod` のあるプロジェクトのルートで実行します。
まずは `localhost:8080` にアクセスしたら `hello` を返す小さなサーバを作ります。

## 目標

- Go の `net/http` パッケージを使う
- `localhost:8080` でサーバを起動する
- ブラウザまたは `curl` でレスポンスを確認する
- ルーティングを少しずつ増やす

## 1. HTTP サーバ用のディレクトリを作る

FizzBuzz 用の `main.go` とは分けて、`api/main.go` を使います。
このリポジトリにはすでにファイルがあるため、作成操作は不要です。
以下はファイルがない状態から始める場合のコマンドです。

```bash
mkdir -p api
touch api/main.go
```

作成後の構成イメージです。

```text
.
├── main.go
├── api/
│   └── main.go
├── calculate/
│   └── fizzbuzz.go
├── go.mod
└── README.md
```

## 2. `api/main.go` にコードを書く

現在のコードは `/` と `/f` で `Hello, World!` を返します。
以下は演習用のコード例です。この手順に沿って進める場合は、内容を次のように書き換えます。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("server started: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
```

## 3. サーバを起動する

```bash
go run ./api
```

ターミナルに次の起動メッセージが表示されます。
直後にエラーが出ず、コマンドが終了せずに待機していることも確認します。

```text
server started: http://localhost:8080
```

このターミナルはサーバ実行中のままにします。
止めたいときは `Ctrl + C` を押します。

## 4. 動作確認する

ブラウザで次の URL を開きます。

```text
http://localhost:8080
```

または、別のターミナルから `curl` で確認します。

```bash
curl http://localhost:8080
```

次のように返ってくれば成功です。

```text
hello
```

## 5. ルートを増やしてみる

`main()` に `/health` 用のルートを追加します。

```go
http.HandleFunc("/health", healthHandler)
```

同じファイルに `healthHandler` を追加します。

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}
```

コードの変更は自動では反映されません。実行中のサーバを `Ctrl + C` で止め、起動し直します。

```bash
go run ./api
```

別のターミナルで確認します。

```bash
curl http://localhost:8080/health
```

```text
ok
```

## 6. 追加課題

ここからは自分で実装してみましょう。`"/"` のハンドラは未登録のパスも受け取るため、
存在しないパスで `404` を返すには、パスを確認する処理などが必要です。

- `/hello` にアクセスしたら `hello` を返す
- `/goodbye` にアクセスしたら `goodbye` を返す
- `/name?value=sakana` にアクセスしたら `hello sakana` を返す
- 存在しないパスでは `404` を返す
- `POST` リクエストだけ受け付ける API を作る

## 完了の目安

- サーバを自分で起動・停止できる
- `/` で `hello`、`/health` で `ok` が返る
- URL とハンドラの関係を説明できる

ポート使用中のエラーが出る場合は、前に起動したサーバが残っていないか確認してください。

[研修1のトップに戻る](../README.md)
