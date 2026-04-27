from fastapi import FastAPI, HTTPException

app = FastAPI(title="計算 API", description="2つの数字を受け取って計算するシンプルな API")


@app.get("/add")
def add(a: float, b: float) -> dict:
    """2つの数字 a と b を受け取り、その合計を返す"""
    return {"a": a, "b": b, "result": a + b}


@app.get("/calculate")
def calculate(a: float, b: float, operator: str = "add") -> dict:
    """2つの数字 a と b を受け取り、operator に合わせて計算する"""
    if operator == "add":
        result = a + b
    elif operator == "subtract":
        result = a - b
    elif operator == "multiply":
        result = a * b
    elif operator == "divide":
        if b == 0:
            raise HTTPException(status_code=400, detail="0で割ることはできません")
        result = a / b
    else:
        raise HTTPException(status_code=400, detail="operator は add, subtract, multiply, divide のどれかを指定してください")

    return {"a": a, "b": b, "operator": operator, "result": result}
