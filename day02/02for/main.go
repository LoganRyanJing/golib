package main

import "fmt"

func main() {
	//常规循环语句
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//变种1，省略初始语句
	// i := 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//变种2，省略结束语句
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range循环
	// s := "Hello大学城"
	// for i, v := range s {
	// 	fmt.Printf("%d %c\n", i, v)
	// }

	// 跳出循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

}
