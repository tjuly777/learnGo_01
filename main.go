package main

import (
	"./animals"
	"fmt"
)

func main() {
	fmt.Println(AppName(2, 5))

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
