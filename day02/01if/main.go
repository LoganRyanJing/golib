package main

import (
	"fmt"
)

func main() {
	// age := 19

	//age变量作用域只在if条件判断语句中生效，减少内存占用
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	} else {
		fmt.Println("该写暑假作业了")
	}

	// fmt.Println(age) //声明在if语句中的变量无法打印

}
