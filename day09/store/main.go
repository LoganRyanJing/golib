package main

import (
	"fmt"
	"store"
)

func main() {
	bruce, _ := store.CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce.CheckCreadits())
	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = bruce.RemoeCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	bruce.ChangeName("Mark")

	fmt.Println(bruce)
}
