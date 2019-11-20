package main

import "fmt"

func main() {
	//十进制
	var i1 = 101
	fmt.Printf("%c\n", 0x4E2D)
	fmt.Printf("%c\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)

	i4 := int8(9)
	fmt.Printf("%T", i4)
}
