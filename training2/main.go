package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

// Trade は1回の取引を表す。同じユーザーの行も、それぞれ1回として数える。
type Trade struct {
	UserID string `json:"user_id"`
}

func loadTrades(path string) ([]Trade, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var trades []Trade
	if err := json.Unmarshal(data, &trades); err != nil {
		return nil, err
	}
	return trades, nil
}

func countTrades(trades []Trade, userID string) int {
	count := 0
	for _, trade := range trades {
		if trade.UserID == userID {
			count++
		}
	}
	return count
}

func newHandler(trades []Trade) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{user_id}/trades", func(w http.ResponseWriter, r *http.Request) {
		userID := r.PathValue("user_id")
		response := struct {
			Count int `json:"count"`
		}{Count: countTrades(trades, userID)}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("レスポンスの書き込みに失敗: %v", err)
		}
	})
	return mux
}

func main() {
	dataPath := flag.String("data", "training2/trades.json", "取引データの JSON ファイル")
	flag.Parse()

	trades, err := loadTrades(*dataPath)
	if err != nil {
		log.Fatalf("取引データの読み込みに失敗: %v", err)
	}

	log.Println("server started: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", newHandler(trades)))
}
