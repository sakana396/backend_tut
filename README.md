# 研修1：Go の基礎と HTTP サーバ入門

Go で小さなプログラムを動かし、簡単な HTTP サーバを作るための入門教材です。
FizzBuzz で関数と条件分岐に触れてから、ブラウザや `curl` でサーバの応答を確認します。

## 学習の進め方

| 順番 | 資料 | 到達目標 |
| --- | --- | --- |
| 0 | [環境構築](docs/setup.md) | Go を実行できる環境を用意する |
| 1 | [FizzBuzz 演習](docs/fizzbuzz.md) | 入力・関数・条件分岐・出力の流れを理解する |
| 2 | [HTTP サーバ演習](docs/http.md) | サーバを起動し、URL ごとの応答を確認する |

Go が使える場合は、FizzBuzz 演習から始めてください。

## まず動かす

ターミナルでこのプロジェクトのディレクトリに移動します。
以降のコマンドは、`go.mod` がある場所で実行してください。

```bash
go version
go run .
```

`15` を入力して Enter を押すと、`FizzBuzz` が表示されます。

HTTP サーバは次のコマンドで起動します。

```bash
go run ./api
```

ブラウザで `http://localhost:8080` を開くと、現在のコードでは `Hello, World!` が表示されます。
サーバを止めるときは、起動したターミナルで `Ctrl + C` を押します。

## ファイル構成

```text
.
├── README.md             # 研修の入口・進め方
├── go.mod                # Go モジュールの設定
├── main.go               # FizzBuzz の入力・出力
├── calculate/
│   └── fizzbuzz.go       # FizzBuzz の処理
├── api/
│   └── main.go           # HTTP サーバ
└── docs/
    ├── setup.md          # 環境構築
    ├── fizzbuzz.md       # FizzBuzz 演習
    └── http.md           # HTTP サーバ演習
```

## 教材の現在の状態

FizzBuzz は実装済みです。HTTP サーバには `/` と `/f` のルートがあり、どちらも `Hello, World!` を返します。
HTTP サーバ演習では、`hello` を返す基本のコードから `/health` の追加へ進みます。
資料内のコード例と、演習中の自分のコードを見比べながら進めてください。

## 次の課題：研修2

[ユーザーごとの取引回数を返す API](docs/trades.md)では、JSON ファイルの読み込み、
パスからのユーザー ID の取得、件数の計算、JSON レスポンスの返却を学びます。

```bash
go run ./training2
```

研修1のサーバを停止してから起動してください。コードとサンプルデータは `training2/` にあります。
