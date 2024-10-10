package main

import "fmt"

func findWordsContaining(words []string, x byte) []int {
	var wordContainingX []int

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if words[i][j] == x {
				wordContainingX = append(wordContainingX, i)
				break
			}
		}
	}
	return wordContainingX
}

func quadOrdering(arr []int) []int {
	result := make([]int, len(arr))
	i, j, pos := 0, len(arr)-1, len(arr)-1

	for i <= j {
		if arr[i]*arr[i] > arr[j]*arr[j] {
			result[pos] = arr[i] * arr[i]
			i++
		} else {
			result[pos] = arr[j] * arr[j]
			j--
		}
		pos--
	}
	return result
}

func main() {
	fmt.Println(findWordsContaining([]string{"leetcode", "stack", "binary tree"}, 'e'))
	fmt.Println(quadOrdering([]int{-5, -3, 2, 4}))
}
