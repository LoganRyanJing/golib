package main

import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		fizzBuzz(i)
// 	}
// }

// func fizzBuzz(i int) {
// 	fizz := "fizz"
// 	buzz := "buzz"
// 	fizzbuzz := "fizzBuzz"

// 	switch i {
// 	case i%3 == 0 && i%5 == 0:
// 		fmt.Println(i, fizzbuzz)
// 	case i%3 == 0:
// 		fmt.Println(i, fizz)
// 	case i%5 == 0:
// 		fmt.Println(i, buzz)
// 	default:
// 		fmt.Println(i)
// 	}

// 	return
// }

// 在 Go 语言中，switch 语句的 case 表达式必须是常量或者常量表达式。
// i%3 == 0 && i%5 == 0 是一个布尔表达式，不是常量或常量表达式，所以不能用在 switch 语句的 case 中。
// 你可以将 i%3 == 0 && i%5 == 0 放在 if 语句中，然后在 if 语句中调用 fizzBuzz 函数。这样就可以避免 switch 语句的限制。

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fizzBuzz(i, "fizzBuzz")
// 		} else if i%3 == 0 {
// 			fizzBuzz(i, "fizz")
// 		} else if i%5 == 0 {
// 			fizzBuzz(i, "buzz")
// 		} else {
// 			fizzBuzz(i, "")
// 		}
// 	}
// }

// func fizzBuzz(i int, str string) {
// 	fmt.Println(i, str)
// }

// func fizzbuzz(num int) string {
// 	switch {
// 	case num%15 == 0:
// 		return "FizzBuzz"
// 	case num%3 == 0:
// 		return "Fizz"
// 	case num%5 == 0:
// 		return "Buzz"
// 	}
// 	return strconv.Itoa(num)
// }

// func main() {
// 	for number := 1; number <= 100; number++ {
// 		fmt.Println(fizzbuzz(number))
// 	}
// }

// 查找质数
// func findprimes(number int) bool {
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			return false
// 		}
// 	}

// 	if number > 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	fmt.Println("Prime number less than 20")

// 	for number := 0; number <= 20; number++ {
// 		if findprimes(number) {
// 			fmt.Printf("%v ", number)
// 		}
// 	}
// }

func main() {
	val := 0

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
