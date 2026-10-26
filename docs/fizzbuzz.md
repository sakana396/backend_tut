# FizzBuzz 演習

[研修1のトップに戻る](../README.md)

整数を1つ受け取り、条件に応じた文字列を返す関数を扱います。
現在の `calculate/fizzbuzz.go` には実装が入っています。まず実行して動きを確認し、条件分岐を読んでみましょう。

すべてのコマンドは、`go.mod` のあるプロジェクトのルートで実行します。

## プロジェクト構成

```
.
├── main.go              # エントリーポイント（編集不要）
├── calculate/
│   └── fizzbuzz.go      # FizzBuzz ロジック（演習対象）
├── go.mod
└── README.md
```

## FizzBuzz のルール

- 3 の倍数 → "Fizz"
- 5 の倍数 → "Buzz"
- 15 の倍数 → "FizzBuzz"
- それ以外 → その数値

## 実行方法

```bash
go run main.go
# 入力例: 15
# 出力: FizzBuzz
```

## テスト例

```bash
# 15 → FizzBuzz
go run main.go <<< "15"

# 3 → Fizz
go run main.go <<< "3"

# 5 → Buzz
go run main.go <<< "5"

# 2 → 2
go run main.go <<< "2"
```

## 学習ポイント

- **Go の基本**: パッケージ構造、エクスポート（大文字）
- **関数実装**: ロジック実装と単純な条件分岐

## 編集対象

`calculate/fizzbuzz.go` の `FizzBuzz()` 関数を演習対象にします。`main.go` の変更は不要です。

## 確認してみよう

- 15 の倍数の判定を先に書いているのはなぜでしょうか。
- `strconv.Itoa` は何をしているでしょうか。
- 3・5・15・2 を入力し、上のルールどおりになることを確認しましょう。

確認できたら、[HTTP サーバ演習](http.md)へ進みます。

