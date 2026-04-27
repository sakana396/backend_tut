# 計算 API チュートリアル

Python の **FastAPI** を使って、2 つの数字を受け取って計算した結果を返す Web API を作るチュートリアルです。

---

## 目次

1. [前提条件](#前提条件)
2. [セットアップ](#セットアップ)
3. [コードの説明](#コードの説明)
4. [サーバーの起動](#サーバーの起動)
5. [API の使い方](#api-の使い方)
6. [テストの実行](#テストの実行)

---

## 前提条件

- Python 3.9 以上がインストールされていること

---

## セットアップ

```bash
# 依存パッケージをインストール
pip install -r requirements.txt
```

---

## コードの説明

`main.py` の内容を確認してみましょう。

```python
from fastapi import FastAPI

app = FastAPI(title="計算 API", description="2つの数字を受け取って計算するシンプルな API")


@app.get("/add")
def add(a: float, b: float) -> dict:
    """2つの数字 a と b を受け取り、その合計を返す"""
    return {"a": a, "b": b, "result": a + b}
```

| 行 | 説明 |
|----|------|
| `FastAPI(...)` | アプリケーションのインスタンスを作成します |
| `@app.get("/add")` | `GET /add` というエンドポイントを定義します |
| `a: float, b: float` | クエリパラメータとして 2 つの数字を受け取ります |
| `return {...}` | 入力値と計算結果を JSON で返します |

---

## サーバーの起動

```bash
uvicorn main:app --reload
```

起動すると以下の URL でアクセスできます。

- API: `http://127.0.0.1:8000`
- 自動生成ドキュメント (Swagger UI): `http://127.0.0.1:8000/docs`

---

## API の使い方

### エンドポイント

```
GET /add?a={数字1}&b={数字2}
GET /calculate?a={数字1}&b={数字2}&operator={計算方法}
```

`operator` には以下を指定できます。

- `add`: 足し算
- `subtract`: 引き算
- `multiply`: 掛け算
- `divide`: 割り算

### リクエスト例

```bash
curl "http://127.0.0.1:8000/add?a=3&b=5"
curl "http://127.0.0.1:8000/calculate?a=10&b=4&operator=subtract"
```

### レスポンス例

```json
{
  "a": 3.0,
  "b": 5.0,
  "result": 8.0
}
```

```json
{
  "a": 10.0,
  "b": 4.0,
  "operator": "subtract",
  "result": 6.0
}
```

ブラウザで `http://127.0.0.1:8000/docs` を開くと、GUI 上で API を試すこともできます。

---

## テストの実行

```bash
pytest test_main.py -v
```

テストが全て通れば実装は完了です 🎉
