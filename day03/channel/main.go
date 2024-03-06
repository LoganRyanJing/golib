package main

import "fmt"

func worker(c chan int) {
	fmt.Println("I am worker...")

	num := <-c
	fmt.Println("得到从main传递的数据是：", num)
}

func main() {
	// 创建一个channel
	c := make(chan int)

	//开辟一个协程去执行worker
	go worker(c)

	// main向c中传入一个数据
	c <- 2
	fmt.Println("I am main...")
}
