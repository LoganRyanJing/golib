package main

import "fmt"

func main() {
	firstName := "John"
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	*name = "David"
}

//  & 运算符使用其后对象的地址。
//  * 运算符取消引用指针。 你可以前往指针中包含的地址访问其中的对象。
