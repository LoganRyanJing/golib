// Go 中的数组是一种特定类型且长度固定的数据结构。 它们可具有零个或多个元素，你必须在声明或初始化它们时定义大小。
// 此外，它们一旦创建，就无法调整大小。 鉴于这些原因，数组在 Go 程序中并不常用，但它们是切片和映射的基础。
package main

import (
	"fmt"
)

// func main() {
// 	var a [3]int

// 	a[1] = 10
// 	fmt.Println(a[0])
// 	fmt.Println(a[1])
// 	fmt.Println(a[len(a)-1])
// }

func main() {
	// cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	// fmt.Println("Cities:", cities)
	// fmt.Println(len(cities))

	// numbers := [...]int{99: -1, 1: 200}
	// fmt.Println("First Position:", numbers[0])
	// fmt.Println("Last Position:", numbers[99])
	// fmt.Println("Length:", len(numbers))
	// fmt.Println(numbers)

	//二维数组
	// var twoD [3][5]int
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		twoD[i][j] = (i + 1) * (j + 1)
	// 	}
	// 	fmt.Println("Row", i, twoD[i])
	// }
	// fmt.Println("\nAll at once:", twoD)

	// 三维数组
	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once：", threeD)
}
