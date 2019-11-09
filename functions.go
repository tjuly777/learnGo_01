/**
3.11「関数」で使用した関数郡
関数・戻り値・無名関数・クロージャ・クロージャ内のローカル変数の挙動
*/
package main

import (
	"errors"
	"fmt"
)

func AppName(x, y int) int {
	return getAppName(x, y)
}

func getAppName(x, y int) int {
	return x + y
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
	err = errors.New("This is error.") //interface{}型の初期値、nilではなくなったことでエラーを表す
	return res, err
}

//無名関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function.")
	}
}

//関数を引数に取る関数
func callFunction(f func()) {
	f()
}

//無名関数を返す関数
func later() func(string) string {
	/**
	later()中のローカル関数。通常は関数実行後に破棄されるが、
	クロージャからの参照があった場合は破棄されない。
	*/
	var store string

	//later()実行時に返却されるクロージャ。
	return func(next string) string {
		//呼び出されたタイミングで、前回呼び出した際に保存した引数を返却用変数に代入。
		s := store
		//引数をlater()のローカル変数に保存。
		store = next
		return s
	}
}

//クロージャを応用したジェネレータの習作
func incrementFunc() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
