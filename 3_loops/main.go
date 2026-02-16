package main

import "fmt"

func main() {
	var num int8 = 5

	// FOR loop

	for i := range 5 {
		fmt.Println(i + 1)
	}

	// While loops

	for num > 0 {
		fmt.Println(num)
		num -= 1
	}

	var nums []int

	for i := range 5 {
		nums = append(nums, i+1)
	}

	fmt.Println((nums))
}
