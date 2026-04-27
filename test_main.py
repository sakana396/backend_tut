from fastapi.testclient import TestClient

from main import app

client = TestClient(app)


def test_add_integers():
    response = client.get("/add", params={"a": 3, "b": 5})
    assert response.status_code == 200
    assert response.json() == {"a": 3.0, "b": 5.0, "result": 8.0}


def test_add_floats():
    response = client.get("/add", params={"a": 1.5, "b": 2.5})
    assert response.status_code == 200
    assert response.json() == {"a": 1.5, "b": 2.5, "result": 4.0}


def test_add_negative_numbers():
    response = client.get("/add", params={"a": -3, "b": 7})
    assert response.status_code == 200
    assert response.json() == {"a": -3.0, "b": 7.0, "result": 4.0}


def test_add_missing_parameter():
    response = client.get("/add", params={"a": 3})
    assert response.status_code == 422


def test_calculate_add():
    response = client.get("/calculate", params={"a": 3, "b": 5, "operator": "add"})
    assert response.status_code == 200
    assert response.json() == {"a": 3.0, "b": 5.0, "operator": "add", "result": 8.0}


def test_calculate_subtract():
    response = client.get("/calculate", params={"a": 10, "b": 4, "operator": "subtract"})
    assert response.status_code == 200
    assert response.json() == {"a": 10.0, "b": 4.0, "operator": "subtract", "result": 6.0}


def test_calculate_multiply():
    response = client.get("/calculate", params={"a": 3, "b": 5, "operator": "multiply"})
    assert response.status_code == 200
    assert response.json() == {"a": 3.0, "b": 5.0, "operator": "multiply", "result": 15.0}


def test_calculate_divide():
    response = client.get("/calculate", params={"a": 10, "b": 2, "operator": "divide"})
    assert response.status_code == 200
    assert response.json() == {"a": 10.0, "b": 2.0, "operator": "divide", "result": 5.0}


def test_calculate_divide_by_zero():
    response = client.get("/calculate", params={"a": 10, "b": 0, "operator": "divide"})
    assert response.status_code == 400


def test_calculate_invalid_operator():
    response = client.get("/calculate", params={"a": 10, "b": 2, "operator": "power"})
    assert response.status_code == 400
