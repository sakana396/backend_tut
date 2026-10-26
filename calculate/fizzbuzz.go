package calculate

import "strconv"

// FizzBuzz は与えられた数値に対して FizzBuzz ルールを適用した結果を返す
// - 15 の倍数: "FizzBuzz"
// - 3 の倍数: "Fizz"
// - 5 の倍数: "Buzz"
// - その他: 数値を文字列化した値
func FizzBuzz(number int) string {
	if number%15 == 0 {
		return "FizzBuzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}
