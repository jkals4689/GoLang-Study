package main

import "fmt"

var g int = 10 // global variable

func main() {
	m := 20

	{ // 지역 변수
		s := 50
		fmt.Println(m, s, g)
	} // 지역 함수 삭제

	// m = s + 20 (err. s는 지역변수 안에 있었고 지역 함수가 없어짐에 따라 s도 사라짐.)
}
