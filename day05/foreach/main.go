package main

import (
	"fmt"
)

// func main() {
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 	}

// 	fmt.Println("sum of 1..100 is", sum)
// }

// Go没有while关键字，但可以改用 for 循环，
// 并利用 Go 使预处理语句和后处理语句可选这一事实。

// func main() {
// 	var num int64
// 	rand.Seed(time.Now().Unix())
// 	for num != 5 {
// 		num = rand.Int63n(15)
// 		fmt.Println(num)
// 	}
// }

// func main() {
// 	var num int32
// 	sec := time.Now().Unix()
// 	rand.Seed(sec)

// 	for {
// 		fmt.Print("Writing inside the loop...")
// 		if num = rand.Int31n(10); num == 5 {
// 			fmt.Println("finish")
// 			break
// 		}
// 		fmt.Println(num)
// 	}
// }

func main() {
	sum := 0
	for num := 1; num < 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers dicisible by 5, is", sum)
}
