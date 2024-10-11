package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int)

	for _, num := range nums {
		if _, ok := hashMap[num]; ok {
			return true
		}
		hashMap[num] = num
	}
	return false
}

func main() {
	fmt.Println("[false]:", containsDuplicate([]int{1, 4, 3, 6, 7}))
	fmt.Println("[true]: ", containsDuplicate([]int{1, 4, 3, 3, 6, 7}))
}
