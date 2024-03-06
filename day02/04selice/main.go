package main

import "fmt"

// 切片selice
func main() {
	// 切片的定义
	var s1 []int
	var s2 []string

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组产生切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s3 := a1[0:4] //对数组进行切割得到切片，左包含右不包含（左闭右开）
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4]
	s6 := a1[3:]
	s7 := a1[:]
	fmt.Println(s5, s6, s7)
	// 切片的容量是从底层数组的第一个元素指向最后一个元素的容量，支持扩容
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))

	//切片再切割
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8))
	// 切片是引用类型，都指向底层数组
	fmt.Println("s6:", s6)
	a1[6] = 1300
	fmt.Println("s6:", s6)

	// 切片的遍历
	// 1.索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// 2.for range循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
