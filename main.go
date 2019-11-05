package main

import (
	"./animals"
	"errors"
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
		fmt.Printf("Result == %v. Error Cause '%v'\n", result, err)
	}

	//変数に無名関数を代入することも可能。
	f := func(x, y int) int { return x + y }
	fmt.Println(f(1, 11))

	//関数を返す関数
	f2 := returnFunc()
	f2()
	//変数を経由せずに直接呼び出すことも可能。
	returnFunc()()

	//関数を引数としてとる。ここでは、無名関数を引数として与えている。
	callFunction(func() {
		fmt.Println("Called Function.")
	})

	//クロージャから関数内のローカル変数にアクセスした際の挙動確認
	f4 := later()
	fmt.Println(f4("It"))
	fmt.Println(f4("is"))
	fmt.Println(f4("Closure Lesson."))

	//上記を応用した、ジェネレータ
	f5 := incrementFunc()
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())

	/**
	新たにクロージャを生成すると、値は別個となる。
	オブジェクトのインスタンス化に近いイメージ！
	*/
	f5_1 := incrementFunc()
	fmt.Println(f5_1())
	fmt.Println(f5())

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
