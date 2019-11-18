/**
3.14「制御構文」の学習範囲

*/
package main

import "fmt"

func main() {
	//型アサーション
	getAssertion(10)

	//ラベル
	//defer
	//goroutin
}

/**
型アサーションの実践。
interface{}型の引数において、型を動的に判断する
*/
func getAssertion(a interface{}) {
	_, isInt := a.(int) //型アサーション。tには値、isIntにはboolが入る。
	if isInt == false {
		fmt.Println("Assertion: a is Int.")
	} else {
		fmt.Println("Assertion: a is not Int.")

	}
}
