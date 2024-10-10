package main

import (
	"fmt"
	"unicode"
)

func checkPalindrome(word string) bool {
	var filtered []rune

	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}

	for i := 0; i < len(filtered)/2; i++ {
		if filtered[i] != filtered[len(filtered)-i-1] {
			return false
		}
	}
	return true
	//left, right := 0, len(filtered)-1
	//
	//for left < right {
	//	if filtered[left] != filtered[right] {
	//		return false
	//	}
	//	left++
	//	right--
	//}
	//return true
}

func main() {
	fmt.Println(checkPalindrome("A man, a plan, a canal, Panama"))
}
