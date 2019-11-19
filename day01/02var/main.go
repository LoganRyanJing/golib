package main

import "fmt"

//声明变量
// var name string
// var age int
// var isOk bool

//批量声明变量
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "Logan"
	age = 23
	isOk = true
	fmt.Println("name is %s", name)
	fmt.Println(isOk)
	fmt.Print(age)

	var s1 string = "web"
	fmt.Println(s1)
	var s2 = "haha"
	fmt.Println(s2)
	s3 := "哈哈哈"
	fmt.Println(s3)

	// s1 := "abbdbdd" //同一个作用域不能申明同名变量
}
