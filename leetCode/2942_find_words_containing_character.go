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

func main() {
	fmt.Println(findWordsContaining([]string{"leetcode", "stack", "binary tree"}, 'e'))
}
