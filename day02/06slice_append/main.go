package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s1[3] = "广州" //错误的写法，会导致编译错误：索引越界
	fmt.Println(s1)
	fmt.Printf("\n")

	// 调用append函数一般用原来的切片变量接收返回值
	// append追加元素，原来底层数组放不下，Go底层就会换一个新的底层数组
	// 必须用变量接收append的返回值
	// s1[3] = "广州" //错误的写法，会导致编译错误：索引越界
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) // ...表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
