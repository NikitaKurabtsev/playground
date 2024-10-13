package main

import (
	"errors"
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

func concatStringsBad(string1, string2 string) (string, error) {
	var resString string

	if string1 == "" {
		return "", errors.New("string1 is empty")
	} else {
		if string2 == "" {
			return "", errors.New("string2 is empty")
		} else {
			resString = string1 + string2
		}
	}
	return resString, nil
}

func concatStringsGood(s1, s2 string, maxLength int) (string, error) {
	if s1 == "" {
		return "", errors.New("string1 is empty")
	}
	if s2 == "" {
		return "", errors.New("string2 is empty")
	}

	resString, err := concatString(s1, s2)
	if err != nil {
		return "", err
	}

	if len(resString) > maxLength {
		return resString[:maxLength], nil
	}

	return resString, nil

}

func concatString(s1, s2 string) (string, error) {
	if len(s1)+len(s2) < len(s1) || len(s1)+len(s2) < len(s2) {
		return "", errors.New("concatenation overflow")
	}
	return s1 + s2, nil
}

func main() {
	fmt.Println(checkPalindrome("A man, a plan, a canal, Panama"))

}
