package main

import (
	"fmt"
	"os"
	"strconv"
)

// 如果是创建Go包，则无需main函数
// func main() {
// 	number1, _ := strconv.Atoi(os.Args[1])
// 	number2, _ := strconv.Atoi(os.Args[2])
// 	fmt.Println("Sum:", number1+number2)
// }

// func main() {
// 	sum := sum(os.Args[1], os.Args[2])
// 	fmt.Println("Sum:", sum)
// }

// func sum(number1 string, number2 string) int {
// 	int1, _ := strconv.Atoi(number1)
// 	int2, _ := strconv.Atoi(number2)
// 	return int1 + int2
// }

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1) // 如果不需要函数某个返回值，可以通过_变量放弃该函数
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func main() {
	sum, _ := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	// fmt.Println("Mul:", mul)
}
