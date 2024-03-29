package main

import "fmt"

// copy
func main() {
	// a1 := []int{1, 3, 5}
	// a2 := a1
	// a3 := make([]int, 3, 3)
	// copy(a3, a1) //copy
	// fmt.Println(a1, a2, a3)
	// a1[0] = 100
	// fmt.Println(a1, a2, a3)

	// a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// a = append(a2[:2], a[3:]...)
	// fmt.Println(a)
	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))
	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &s1[0])
	s1 = append(s1[:1], s1[2:]...)
	fmt.Printf("%p\n", &s1[0])
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1)
}
