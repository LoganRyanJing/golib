package main

import "fmt"

func main() {
	const (
		// 从第一行开始，iota 从 0 逐行加一
		a1 = 1 //0
		b1
		a2
		a3 = iota
		a4 = "hello"
		a5 = "hello"
		a6 = iota
		a7 = iota + 1
	)

	fmt.Println(a1, b1, a2, a3, a4, a5, a6, a7)
}
