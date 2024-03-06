package main

import "fmt"

//swich 简化大量的判断
func main() {

	switch n := 3; n {
	case 1:
		fmt.Println("大哥")
	case 2:
		fmt.Println("二哥")
	case 3:
		fmt.Println("老妖")
	}

}
