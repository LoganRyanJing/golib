package main

import "fmt"

func givemenumber() int {
	return -7
}

func main() {
	num := givemenumber()
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digt")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println(num)
}
