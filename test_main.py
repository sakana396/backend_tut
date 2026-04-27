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
