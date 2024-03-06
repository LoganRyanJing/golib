package main

import (
	"fmt"
)

// func main() {
// 	studentsAge := map[string]int{
// 		"John": 32,
// 		"Bob":  31,
// 	}

// 	fmt.Println(studentsAge)
// }

// func main() {
// 	studentsAge := make(map[string]int)
// 	studentsAge["john"] = 32
// 	studentsAge["bob"] = 31
// 	fmt.Println(studentsAge)
// }

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	// fmt.Println("Bob's age is", studentsAge["bob"])
	// fmt.Println("Christy's age is", studentsAge["christy"])

	// 在很多情况下，访问映射中没有的项时 Go 不会返回错误，这是正常的。
	// 但有时需要知道某个项是否存在。 在 Go 中，映射的下标表示法可生成两个值。
	// 第一个是项的值。 第二个是指示键是否存在的布尔型标志。
	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}
}
