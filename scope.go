/**
3.13「スコープ」の学習範囲
パッケージ内の関数などの公開範囲の確認
*/
package main

import (
	f "fmt"  //インポートパッケージの名前を上書きできる（≠エイリアス）
	. "math" //パッケージ名を省略することもできる。パッケージ内の識別子が被らないように注意
)

//パッケージに定義された定数・変数・関数は、識別子が大文字なら他のパッケージからも参照可能
const A = 1

var Example = 512

func main() {
	f.Println("Rewrite Package name.")
	f.Printf("Abbreviation Package name. Pi = %v \n", Pi)
	f.Println()
}
