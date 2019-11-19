/**
3.14「制御構文」の学習範囲

*/
package main

import "fmt"

func main() {
	//型アサーション
	getAssertion(10)
	getAssertionWithSwitch("String")

	//ラベル
	//defer
	//goroutin
}

/**
型アサーションの実践。
interface{}型の引数において、型を動的に判断する
*/
func getAssertion(a interface{}) {
	//この時点では、aの元の型は失われているので・・・
	_, isInt := a.(int) //型アサーションでaの型を復元。tには値、isIntにはboolが入る。
	if isInt == false {
		fmt.Println("Assertion: a is Int.")
	} else {
		fmt.Println("Assertion: a is not Int.")
	}
}

/**
switch文と組み合わせた型アサーションの実践。
*/
func getAssertionWithSwitch(a interface{}) {
	//下記の書き方で、aの型によるswitch分岐が可能。多岐にわたる分岐の場合は、こちらがわかりやすい。
	switch a.(type) {
	case int:
		fmt.Println("Assertion with Switch: a is Int.")
	case string:
		fmt.Println("Assertion with Switch: a is String.")
	case bool:
		fmt.Println("Assertion with Switch: a is Boolean.")
	default:
		fmt.Println("Assertion with Switch: a is the Other type.")
	}
}
