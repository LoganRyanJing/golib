// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// }
package main

import (
	"fmt"

	"github.com/loganryanjing/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version:", calculator.Version)
}
