# 環境構築

[研修1のトップに戻る](../README.md)

## 用意するもの

- ターミナル
- テキストエディタ
- Go（`go.mod` の指定は 1.22 以上。インストール時は公式サイトの現行安定版を使用）
- ブラウザ、または `curl`（HTTP サーバの確認用）

## 1. Windows で WSL を用意する

Windows では WSL の Ubuntu を使う想定です。macOS・Linux の場合は次へ進んでください。

管理者として PowerShell を開いて実行します。

```powershell
wsl --install
```

PC を再起動し、Ubuntu の初回起動時にユーザー名とパスワードを設定します。
以降の Go のインストールと演習は、Ubuntu のターミナルで行います。

対応する Windows の条件や、すでに WSL がある場合の手順は、
[Microsoft の WSL インストールガイド](https://learn.microsoft.com/ja-jp/windows/wsl/install)で確認してください。

## 2. Go をインストールする

[Go 公式のインストールガイド](https://go.dev/doc/install)を開き、使用する OS の手順に従ってください。
WSL の場合は Linux 向けの手順を使います。
ダウンロードするファイルは、OS と CPU の種類に合うものを選んでください。

Linux のアーカイブを使う場合は、ダウンロードしたファイル名と展開コマンドのファイル名を一致させます。
バージョンごとにファイル名が変わるため、公式ガイドで確認してください。

インストール後、新しいターミナルで確認します。

```bash
go version
```

Go のバージョンと OS・CPU の情報が表示されれば準備完了です。
`go: command not found` と表示される場合は、インストール先と `PATH` の設定を公式ガイドで確認してください。

## 3. 教材を実行する

ターミナルで、この教材を配置したディレクトリに移動します。
以下のパスは、自分の配置先に置き換えてください。

```bash
cd /path/to/backend_tut
go run .
```

`15` を入力して Enter を押し、`FizzBuzz` が表示されることを確認します。

次は [FizzBuzz 演習](fizzbuzz.md)へ進みます。
