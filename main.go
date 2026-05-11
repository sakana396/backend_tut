package main

import (
	"calculate"
	"fmt"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	num := calculate.FizzBuzz(input)

	fmt.Println(num)
}
