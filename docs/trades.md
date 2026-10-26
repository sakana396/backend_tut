# 研修2：ユーザーごとの取引回数を返す API

[教材のトップに戻る](../README.md)

## 目標

研修1で作った HTTP サーバを土台に、URL からユーザー ID を受け取り、
そのユーザーの取引回数を JSON で返せるようにします。

- パスから値を取得する
- JSON ファイルを読み込み、構造体とスライスでデータを扱う
- 条件に合うデータを数える
- HTTP レスポンスを JSON にする

## API の仕様

```text
GET /{user_id}/trades
```

`{user_id}` は調べたいユーザーの ID に置き換えます。

| レスポンス項目 | 型 | 内容 |
| --- | --- | --- |
| `count` | 整数 | 指定されたユーザーの取引回数 |

成功時は `200 OK` と `Content-Type: application/json` を返します。

```shell
curl 'http://localhost:8080/A1B2C3D4E5/trades'
```

```json
{
  "count": 197
}
```

ユーザー ID の形式や、取引がない場合の扱いは提示された仕様では未指定です。
この教材では、ID は空でないパス1区間として扱い、取引がないユーザーは `{"count":0}` を返します。
別の URL は `404`、POST などの非対応メソッドは `405` とします。
Go のルーティングの仕様上、GET と同じ URL の HEAD も受け付けます。

## 実装する順序

1. JSON ファイルを読み、取引データのスライスに変換する。
2. `GET /{user_id}/trades` のルートを登録する。
3. `r.PathValue("user_id")` で ID を取得する。
4. 取引データを順に調べ、その ID と一致する件数を数える。
5. `count` を持つレスポンス用の構造体を作る。
6. `Content-Type` を設定し、`json.NewEncoder(w).Encode(...)` で返す。
7. 複数のユーザーや取引がないユーザーでも確認する。

`count` を常に `197` にするのではなく、対象ユーザーのデータから計算しましょう。

## 取引データ

`training2/trades.json` は研修用の架空データです。配列の1要素を1回の取引として数えます。
`A1B2C3D4E5` の取引が197件、`B1B2C3D4E5` の取引が3件入っています。
同じ `user_id` の要素が複数あるのは、それぞれ別の取引を表すためです。

形式の例です。この例では `user-a` が2回、`user-b` が1回になります。

```json
[
  {"user_id": "user-a"},
  {"user_id": "user-b"},
  {"user_id": "user-a"}
]
```

## 実装例を動かす

`training2/main.go` に動作する実装例を用意しています。
自力で取り組む場合は、先に API の仕様と実装する順序を読んでから、自分でコードを書いてください。

研修1のサーバが動いている場合は `Ctrl + C` で停止します。
プロジェクトのルート（`go.mod` のある場所）から起動します。

```bash
go run ./training2
```

別のターミナルで確認します。

```bash
curl -i 'http://localhost:8080/A1B2C3D4E5/trades'
# 200 / {"count":197}

curl -i 'http://localhost:8080/B1B2C3D4E5/trades'
# 200 / {"count":3}

curl -i 'http://localhost:8080/unknown/trades'
# 200 / {"count":0}

curl -i -X POST 'http://localhost:8080/A1B2C3D4E5/trades'
# 405

curl -i 'http://localhost:8080/A1B2C3D4E5/other'
# 404
```

別の JSON ファイルを使う場合は、パスを指定します。

```bash
go run ./training2 -data /path/to/trades.json
```

ファイルは起動時に読み込みます。内容を編集したらサーバを再起動してください。
ファイルがない場合や JSON が壊れている場合は、エラーを表示して終了します。
この演習では、データは上記の形式に従い、各要素に空でない `user_id` があることを前提にします。

## 完了の目安

- URL のユーザー ID を変えると、対応する件数が返る
- 他のユーザーの取引が件数に混ざらない
- 取引がない場合も、文字列ではなく数値の `0` が返る
- JSON に取引を1件追加して再起動すると、そのユーザーの件数が1増える
- ファイル読み込み・件数の計算・JSON の返却の流れを説明できる

確認用テストは次のコマンドで実行できます。

```bash
go test ./training2
```

## 参考資料

- [Go 公式：ServeMux のルーティング](https://pkg.go.dev/net/http#ServeMux)
- [Go 公式：Request.PathValue](https://pkg.go.dev/net/http#Request.PathValue)
- [Go 公式：JSON エンコーダ](https://pkg.go.dev/encoding/json#Encoder.Encode)
