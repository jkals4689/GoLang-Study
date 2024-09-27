package main

import "fmt"

func main() {
	a := 3              // int
	var b float32 = 3.5 // float32

	var c int = int(b)  // float32 -> int Type change.
	d := float64(a * c) // int -> float64

	var e int64 = 7
	f := int64(d) * e // float64 -> int64

	var g int = int(b * 3) // float64 -> int
	var h int = int(b) * 3 // float32 -> int. not eq g.

	fmt.Println(g, h, f)
}
