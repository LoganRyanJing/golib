// 多路复用
// select 语句的工作方式类似于 switch 语句，但它适用于 channel。
//它会阻止程序的执行，直到它收到要处理的事件。 如果它收到多个事件，则会随机选择一个
package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}
