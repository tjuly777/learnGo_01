/**
「4 参照型」の学習範囲。
*/
package main

import "fmt"

func main() {
	//スライス
	getSlice(4, 6)

	//マップ
	//チャネル
}

/**
スライスの挙動確認
宣言時に指定する各要素
	・要素数	=> 初期化する要素の数
	・容量	=> 要素の最大収容数
*/
func getSlice(length int, capacity int) {
	i := 0
	s := make([]int, length, capacity) //要素数、容量を指定してスライスを生成
	//len()、cap()でそれぞれ要素数・容量を取得できる
	fmt.Printf("Slice:var s's lentgh = %v, capacity = %v \n", len(s), cap(s))
	//配列同様、rangeでループすることが可能（要素数分）
	for i, _ = range s {
		s[i] = i
	}
	/**
	スライスの一部を元にしたスライスの生成も可能。容量は引き継がれる。
	下記は、sの要素に対して、始点:終点を指定。指定しないことも可能
	*/
	s2 := s[2:]
	//appendでスライスの末尾に追記ができる
	s2 = append(s2, 9, 9, 9, 9, 10)
	fmt.Println(s2)
	//TODO::参照についての理解を！
}
