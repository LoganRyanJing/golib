package main

import "fmt"
import "time"

func go_worker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我是go的一个协程，我的名字是", name)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("over")
}

func main() {
	go go_worker("小黑")
	go go_worker("小白")

	for i := 0; i < 10; i++ {
		fmt.Println("我是main...")
		time.Sleep(1 * time.Second)

	}
}
