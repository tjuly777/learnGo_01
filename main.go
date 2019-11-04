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

	//複数の戻り値を取得
	q1, r1 := getMultiReturn(7, 3)
	fmt.Printf("Quotient = %v, Remainder = %v\n", q1, r1)

	//_を使えば、戻り値の一部を破棄できる
	q2, _ := getMultiReturn(7, 3)
	fmt.Printf("Quotient = %v, Remainder = null\n", q2)

	/**
	エラー処理の慣例的な書き方。複数の戻り地の一部に、本来nilを採る変数errを充てる。
	errがnilでないことで、エラーかどうかを検知する。
	*/
	result, err := getSomeError()
	if err != nil {
		fmt.Printf("Result == %v. Error Cause '%v'", result, err)
	}
}

//複数の戻り値を返す関数
func getMultiReturn(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

//エラーを返す
func getSomeError() (res bool, err interface{}) {
	//戻り値で変数初期化まで行っているので、下記は省略可！
	//var (
	//	res bool
	//	err interface{}
	//)
	err = "Not nil" //interface{}型の初期値、nilではなくなったことでエラーを表す
	return res, err
}
