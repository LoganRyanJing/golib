package main

import "fmt"

func main() {
	f1 := 1.2344567
	fmt.Printf("%T\n", f1) //默认float类型是64位
	f2 := float32(0.1234566)
	fmt.Printf("%T\n", f2) //显示声明float32
}
