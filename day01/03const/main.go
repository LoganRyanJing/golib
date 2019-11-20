package main

//常量定义之后不能修改
// const pi = 3.1415

// //批量声明变量

// const(
// 	status = 200
// 	notfount = 404
// )

//批量声明变量时，如果某一行声明后没有赋值，默认就和上一行一样
const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota //0
	a2
	a3 = iota
)

//插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
)

func main() {
	// println(n1, n2, n3)
	// println(a1, a2, a3)
	// println(c1, c2, c3, c4)

	// println(d1, d2, d3, d4)

	println(KB)
	println("MB = ", MB)
	// println('')
	// println()
	// println()
	// println()
}
