// Go 是一种强类型语言。 你声明的每个变量都绑定到特定的数据类型，并且只接受与此类型匹配的值。

// Go 有四类数据类型：

// 基本类型：数字、字符串和布尔值
// 聚合类型：数组和结构 ?? 切片
// 引用类型：指针、切片、映射、函数和通道
// 接口类型：接口
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var integer32 int = 2147483648
	// var integer32 int = -10
	// fmt.Println(integer32)

	// var float32 float32 = 2147483647
	// var float64 float64 = 9223372036854775807
	// fmt.Println(float32, float64)

	// fmt.Println(math.MaxFloat32, math.MaxFloat64)

	// const e = 2.71828
	// const Avogadro = 6.0221429e23
	// const Planck = 6.62606957e-34

	// fmt.Println(e, Avogadro, Planck)

	// var featureFlag bool = true
	// // 布尔值只有2种表示true和false

	// fmt.Println(featureFlag)

	// var firstName string = "John"
	// lastName := "Deo"

	// fmt.Print(firstName, "\n", lastName, "\r")

	// fullName := "John Doe \t(alias \"Foo\")\n"
	// fmt.Println(fullName)

	// var str float32
	// fmt.Println(str)
	// var defaultInt int
	// var defaultFloat32 float32
	// var defaultFloat64 float64
	// var defaultBool bool
	// var defaultString string
	// fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// var integer16 = 127
	// var integer32 = 32767

	// fmt.Println(int32(integer16) + int32(integer32))

	i, _ := strconv.Atoi("-42") // _变量值会被忽略
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
