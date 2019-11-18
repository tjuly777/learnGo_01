package main

import (
	"./animals" //他パッケージの読み込み
	"fmt"
)

func main() {
	//同パッケージの関数
	fmt.Println(AppName(2, 5))

	//他パッケージの関数
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())

}
