# FizzBuzz - Go版

シンプルな FizzBuzz 実装プロジェクト（onboard 教材用）

## プロジェクト構成

```
.
├── main.go              # エントリーポイント（編集不要）
├── calculate/
│   └── fizzbuzz.go      # FizzBuzz ロジック（ここを編集）
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

`calculate/fizzbuzz.go` の `FizzBuzz()` 関数のみ編集してください。
