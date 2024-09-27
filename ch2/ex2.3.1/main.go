package main

import "fmt"

func main() {
	var b = 3.1415     // b -> float64로 자동 지정
	c := 365           // c -> int로 자동 지정
	s := "hello world" // s -> string로 자동 지정

	// a := 3  int
	// var b float64 = 3.5 float64

	// var c int = b (err. float64 변수를 int에 대입 불가)
	// d := a*b (err. 다른 타입끼리 계산 불가)

	// var e int64 = 7
	// f := a*e (err. 는 int 타입, e는 int64로 같은 정수이지만, 타입이 달라서 연산 불가)
	// var g int = b*3 (err. 실수가 정수로 자동으로 변환 X)

	fmt.Println(b, c, s) // output : 3.1415 365 hello world
}
