package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"E:\\Code\\GOPATH\\src\\github.com\\\"" //反斜杠转义字符串
	fmt.Println(path)

	//多行字符串
	s2 := `
		世情薄
		人情恶
		与送黄昏花易洛
		`
	fmt.Println(s2)

	//字符串拼接
	name := "理想"
	word := "dsb"

	ss := name + word
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(ss1)

	//字符串切割，如果后一个参数为空，则按字符切割，返回字符串切片
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	str := "study go lang"
	//判断字符串是否存在子字符串
	isSon := strings.Contains(str, "go")
	fmt.Println(isSon)

	// var s1 []string = []string{"1", "2", "3", "4"}
	// //将一系列字符串连接为一个字符串，之间用sep来分隔
	// s2 := strings.Join(s1, ",")
	// fmt.Println(s2) //1,2,3,4

	//是否包含
	fmt.Println(strings.Contains(ss, "dsb"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "dsb"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "sdb"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "d"))

	fmt.Println(strings.Join(ret, "+"))
}
