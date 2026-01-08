package main

import (
	"fmt"
)

func main() {
	nums := [10]int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1}
	isPalindrome(nums)

}

func isPalindrome(nums [10]int) {
	for i := 0; i < len(nums)/2; i++ {

		if nums[i] != nums[len(nums)-1-i] {
			fmt.Println("Не палиндром!")
			return

		}

	}
	fmt.Print("Это палиндром!!")
}
