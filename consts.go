/**
3.12「定数」の学習範囲
*/
package main

import (
	"fmt"
)

const (
	X   = 1           //パッケージ内の定数
	Y                 //値は省略もできる。その場合、直前の定義と同じ値になる
	Z         = X + Y //式の結果を代入できる
	I64 int64 = -1    //明示的に型を指定することが可能
	//基本的に最大値なし、変数への代入時に注意
	N = 999999999999999999999999999999999999999999999999999999
	/**
	定数間の計算には、浮動小数点の丸めは発生しない。
	そのため、精度の高い小数点での計算が可能（前述の通り、変数代入時にはエラーが発生する）
	*/
)

const (
	IO1 = iota //constブロックごとに0から始まるiotaが与えられる。参照するたび1増える
	IO2
	IO3
)

func constFunc() (int, int) {
	//関数内の定数
	const Y = 2
	return X, Y
}

func main() {
	x, y := constFunc()
	fmt.Printf("x=%v y=%v \n", x, y)
	fmt.Printf("IO1=%v IO2=%v IO3=%v \n", IO1, IO2, IO3)
}
