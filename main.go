package main

import (
	"./animals"
	"fmt"
)

func main() {
	fmt.Println(AppName(2, 5))

	q, r := getMultiReturn(7, 3)
	fmt.Printf("Quotient = %v, Remainder = %v\n", q, r)

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}

func getMultiReturn(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
