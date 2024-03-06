package main

import "fmt"

// func main() {
// 	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
// 	fmt.Println(months)
// 	fmt.Println("Length:", len(months))
// 	fmt.Println("Capacity:", cap(months))
// }

func main() {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	// quartre1 := months[0:3]
	// quartre2 := months[3:6]
	// quartre3 := months[6:9]
	// quartre4 := months[9:12]
	// fmt.Println(quartre1, len(quartre1), cap(quartre1))
	// fmt.Println(quartre2, len(quartre2), cap(quartre2))
	// fmt.Println(quartre3, len(quartre3), cap(quartre3))
	// fmt.Println(quartre4, len(quartre4), cap(quartre4))

	// quartre2 := months[3:6]
	// quartre2Extended := quartre2[:4]
	// fmt.Println(quartre2, len(quartre2), cap(quartre2))
	// fmt.Println(quartre2Extended, len(quartre2Extended), cap(quartre2Extended))

	// 追加项
	// var numbers []int
	// for i := 0; i < 10; i++ {
	// 	numbers = append(numbers, i)
	// 	fmt.Println("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	// }

	// 删除项
	letters := []string["A", "B", "C", "D", "E"]
	remove := 2

	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove", letters[remove])

		letters = append(letters[:remove], letters[remove + 1:]...)

		fmt.Println("After", letters)
	}
}
