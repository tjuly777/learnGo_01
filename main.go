package main

import (
	"./animals"
	"fmt"
)

func main() {
	fmt.Println(AppName(2, 5))

	q1, r1 := getMultiReturn(7, 3)
	fmt.Printf("Quotient = %v, Remainder = %v\n", q1, r1)

	q2, _ := getMultiReturn(7, 3) //_を使えば、返り値の一部を破棄できる
	fmt.Printf("Quotient = %v, Remainder = null\n", q2)

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}

func getMultiReturn(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
