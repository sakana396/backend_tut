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

## セットアップガイド

### WSL（Windows Subsystem for Linux）のインストール

Windows 10 以降で Golang 開発環境を構築する場合、WSL の使用を推奨します。

#### 1. PowerShell を管理者権限で開く

Windows キーを押して「PowerShell」と入力し、「Windows PowerShell」を右クリックして「管理者として実行」を選択します。

#### 2. WSL をインストール

```powershell
wsl --install
```

このコマンドで WSL2 と Ubuntu がインストールされます。

#### 3. 再起動とセットアップ

インストール後、PC を再起動します。再起動後、Ubuntu が自動的に起動し、ユーザー名とパスワードを設定するよう求められます。

#### 4. WSL Ubuntu を起動

以降、Windows ターミナルから Ubuntu にアクセスできます。ターミナルを開いて以下を実行：

```bash
wsl
```

### Golang のインストール

WSL または Linux 環境での Golang インストール方法です。

#### 1. 最新の Go をダウンロード

```bash
wget https://go.dev/dlgo1.26.3.linux-amd64.tar.gz
```

（バージョンは最新版に合わせてください。[go.dev/dl](https://go.dev/dl) で確認）

#### 2. 既存の Go を削除（初回は不要）

```bash
sudo rm -rf /usr/local/go
```

#### 3. Go をインストール

```bash
sudo tar -C /usr/local -xzf go1.22.linux-amd64.tar.gz
```

#### 4. パスを設定

`.bashrc` または `.zshrc` に以下を追加：

```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

設定を反映：

```bash
source ~/.bashrc  # bash を使用している場合
# または
source ~/.zshrc   # zsh を使用している場合
```

#### 5. インストール確認

```bash
go version
```

`go version go1.22.x linux/amd64` のような出力が表示されればインストール完了です。

### クイックスタート

環境構築が完了したら：

```bash
# プロジェクトディレクトリに移動
cd /path/to/backend_tut

# 実行
go run main.go
```
