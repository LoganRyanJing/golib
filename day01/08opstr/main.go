package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "hello go"
	//判断字符串s是否包含子串
	r := strings.Contains(s, "go")
	fmt.Println(r) //true

	var s1 []string = []string{"1", "2", "3", "4"}
	//将一系列字符串连接为一个字符串，之间用sep来分隔
	s2 := strings.Join(s1, ",")
	fmt.Println(s2) //1,2,3,4

	//子串sep在字符串s中第一次出现的位置，不存在则返回-1
	n1 := strings.Index(s, "go")
	fmt.Println(n1) //6

	//返回count个s串联的字符串
	s3 := strings.Repeat(s, 2)
	fmt.Println(s3) //hello gohello go

	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
	s4 := strings.Replace(s, "o", "e", -1)
	fmt.Println(s4) //helle ge

	//字符串切割，如果后一个参数为空，则按字符切割，返回字符串切片
	s5 := strings.Split(s, " ")
	fmt.Println(s5) //[hello go]

	//切除字符串两端的某类字符串
	s6 := strings.Trim(s, "o")
	fmt.Println(s6) //hello g

	//去除字符串的空格符，并且按照空格分割返回slice
	s7 := " hello go "
	s8 := strings.Fields(s7)
	fmt.Println(s8) //[hello go]
}
