package main

import (
	"fizzbuzz/calculate"
	"fmt"
)

// main は標準入力から数値を受け取り、FizzBuzz の結果を出力する
func main() {
	//変数宣言
	var number int
	// 標準入力から数値を読み込む
	fmt.Scanf("%d", &number)

	// FizzBuzz ロジックを実行
	result := calculate.FizzBuzz(number)

	// 結果を出力
	fmt.Println(result)
}
