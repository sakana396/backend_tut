package main

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestTradesAPI(t *testing.T) {
	trades, err := loadTrades("trades.json")
	if err != nil {
		t.Fatal(err)
	}
	handler := newHandler(trades)
	for _, tc := range []struct {
		name   string
		method string
		path   string
		status int
		count  int
	}{
		{"sample user", "GET", "/A1B2C3D4E5/trades", 200, 197},
		{"other user", "GET", "/B1B2C3D4E5/trades", 200, 3},
		{"no trades", "GET", "/unknown/trades", 200, 0},
		{"unsupported method", "POST", "/A1B2C3D4E5/trades", 405, 0},
		{"unknown route", "GET", "/A1B2C3D4E5/other", 404, 0},
		{"missing user", "GET", "/trades", 404, 0},
		{"extra segment", "GET", "/a/b/trades", 404, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
			if w.Code != tc.status {
				t.Fatalf("status = %d, want %d", w.Code, tc.status)
			}
			if tc.status != 200 {
				return
			}
			if got := w.Header().Get("Content-Type"); got != "application/json" {
				t.Fatalf("Content-Type = %q", got)
			}
			var body map[string]int
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if count, ok := body["count"]; !ok || count != tc.count || len(body) != 1 {
				t.Fatalf("body = %v, want count %d only", body, tc.count)
			}
		})
	}
}

func TestLoadTrades(t *testing.T) {
	path := filepath.Join(t.TempDir(), "trades.json")
	if _, err := loadTrades(path); err == nil {
		t.Fatal("missing file should fail")
	}
	if err := os.WriteFile(path, []byte("broken json"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := loadTrades(path); err == nil {
		t.Fatal("invalid JSON should fail")
	}
	if err := os.WriteFile(path, []byte(`[{"user_id":"a"},{"user_id":"b"},{"user_id":"a"}]`), 0600); err != nil {
		t.Fatal(err)
	}
	trades, err := loadTrades(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := countTrades(trades, "a"); got != 2 {
		t.Fatalf("count = %d, want 2", got)
	}
	if err := os.WriteFile(path, []byte(`[]`), 0600); err != nil {
		t.Fatal(err)
	}
	trades, err = loadTrades(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := countTrades(trades, "a"); got != 0 {
		t.Fatalf("empty file count = %d, want 0", got)
	}
}
