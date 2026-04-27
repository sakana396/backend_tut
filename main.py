from fastapi import FastAPI

app = FastAPI(title="足し算 API", description="2つの数字を足し算して返すシンプルな API")


@app.get("/add")
def add(a: float, b: float) -> dict:
    """2つの数字 a と b を受け取り、その合計を返す"""
    return {"a": a, "b": b, "result": a + b}
