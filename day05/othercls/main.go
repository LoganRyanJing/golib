package main

import (
	"fmt"
)

// func main() {
// 	// defer 语句会推迟函数（包括任何参数）的运行，直到包含 defer 语句的函数完成
// 	for i := 1; i <= 4; i++ {
// 		defer fmt.Println("deferred", -i)
// 		fmt.Println("regular", i)
// 	}
// }

// func main() {
// 	newfile, error := os.Create("learnGo.txt")
// 	if error != nil {
// 		fmt.Println("Error: Could not create file.")
// 		return
// 	}

// 	defer newfile.Close()

// 	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
// 		fmt.Println("Error : Could not write to file.")
// 		return
// 	}

// 	newfile.Sync()
// }

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than hight")
	}
	defer fmt.Println("Deferred : hightlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")
	highlow(high, low+1)
}

func main() {
	// panic 和 recover 函数的组合是 Go 处理异常的惯用方式。
	// 其他编程语言使用 try/catch 块。 Go 首选此处所述的方法。
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
