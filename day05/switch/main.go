package main

// func main() {
// 	sec := time.Now().Unix()
// 	rand.Seed(sec)
// 	i := rand.Int31n(10)

// 	switch i {
// 	case 0:
// 		fmt.Println("zero...")
// 	case 1:
// 		fmt.Println("one...")
// 	case 2:
// 		fmt.Println("two...")
// 	default:
// 		fmt.Println("no match")
// 	}

// 	fmt.Println("OK")
// }

import (
	"fmt"
)

// func loaction(city string) (string, string) {
// 	var region string
// 	var continent string
// 	switch city {
// 	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
// 		region, continent = "India", "Asia"
// 	case "Lafayette", "Louisville", "Boulder":
// 		region, continent = "Colorado", "USA"
// 	case "Irvine", "Los Angeles", "San Diego":
// 		region, continent = "California", "USA"
// 	default:
// 		region, continent = "Unknown", "Unknown"
// 	}
// 	return region, continent
// }

// func main() {
// 	region, continent := loaction("Irvine")
// 	fmt.Printf("John works in %s, %s\n", region, continent)
// }

// func main() {
// 	switch time.Now().Weekday().String() {
// 	case "Monday", "Thursday", "Friday":
// 		fmt.Println("It's time to learn go")
// 	default:
// 		fmt.Println("weekend")

// 	}
// 	fmt.Println(time.Now().Weekday())
// }

// func main() {
// 	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
// 	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

// 	contact := "foo@bar.com"
// 	switch {
// 	case email.MatchString(contact):
// 		fmt.Println(contact, "is a email")
// 	case phone.MatchString(contact):
// 		fmt.Println(contact, "is a phone number")
// 	default:
// 		fmt.Println(contact, "is not recognized")
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	r := rand.Float64()
// 	switch {
// 	case r > 0.1:
// 		fmt.Println("Common case, 90% of the time")
// 	default:
// 		fmt.Println("10% of the time")
// 	}
// }

// 在某些编程语言中，你会在每个 case 语句末尾写一个 break 关键字。
// 但在 Go 中，当逻辑进入某个 case 时，它会退出 switch 块，除非你显式停止它。
// 若要使逻辑进入到下一个紧邻的 case，请使用 fallthrough 关键字

func main() {
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}
