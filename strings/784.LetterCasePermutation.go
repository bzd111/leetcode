package main

import (
	"fmt"
)

func letterCasePermutation(S string) []string {
	ans := []string{}
	dfs([]rune(S), 0, &ans)
	return ans
}

func dfs(s []rune, i int, ans *[]string) {
	if i == len(s) {
		*ans = append(*ans, string(s))
		return
	}
	// if
	dfs(s, i+1, ans)
	e := s[i]
	if !('0' <= e && e <= '9') {
		// (X || Y) && !(X && Y)
		s[i] ^= 32
		dfs(s, i+1, ans)
		s[i] ^= 32
	}
}

func main() {
	s := "a1b2" // ["a1b2", "a1B2", "A1b2", "A1B2"]
	fmt.Println(letterCasePermutation(s))
	s = "3z4" // ["3z4", "3Z4"]
	fmt.Println(letterCasePermutation(s))
	s = "12345" // ["12345"]
	fmt.Println(letterCasePermutation(s))
}
