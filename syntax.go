/**
3.14「制御構文」の学習範囲

*/
package main

import "fmt"

func main() {
	//range・goto
	forByRange()

	//型アサーション
	getAssertion(10)
	getAssertionWithSwitch("String")

	//defer
	runDefer()

	//goroutin
}

/**
rangeを用いたfor文。
配列・文字列など、既存の複数要素のあるデータ型をループする。
PHPのforeachと似た使用感。
*/
func forByRange() {
	x := [5]string{"A", "B", "C", "D", "E"}
	for i, v := range x {
		if i > 2 {
			//ラベルを用いたgotoでループを抜ける。（※gotoの実践なので、本来の用途ならbreakを使うべき）
			goto Coda
		}
		fmt.Printf("forByRange: [%v] => %v \n", i, v)
	}
Coda:
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

/**
関数終了時に実行されるdeferの実践。
deferは登録順ではなく、最後に登録されたものから実行されていくことに注意。
*/
func runDefer() {
	defer fmt.Println("runDefer:Last")
	defer func() {
		x := [5]int{}
		for i, _ := range x {
			fmt.Printf("runDefer: %v\n", i)
		}
	}()
	fmt.Println("runDefer:1st")
	fmt.Println("runDefer:2nd")
}
