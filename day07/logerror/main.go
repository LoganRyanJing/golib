package main

import (
	"log"
	"os"
)

func main() {
	// log.Print("Hey, I'm a log!")
	// log.Fatal("Hey, I'm an error log!")
	// log.Panic("Hey, I'm an error log!")
	// fmt.Println("Can you see me?")

	// 函数是 log.SetPrefix()。 可使用它向程序的日志消息添加前缀。
	// log.SetPrefix("main():")
	// log.Print("Hey, I'm a log")
	// log.Fatal("Hey, I'm an error log!")

	// 将日志发送到文件
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")
}
