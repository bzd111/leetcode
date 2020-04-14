package main

import (
	"fmt"
	"sort"
)

func shiftingLetters(S string, shifts []int) string {
	c := 0
	s := []rune(S)
	sort.Reverse(sort.IntSlice(shifts))
	for i := len(shifts) - 1; i >= 0; i-- {
		c += shifts[i] % 26
		s[i] = rune(((int(s[i])-int('a')+c)%26 + int('a')))
	}

	return string(s)
}

func main() {
	// Input: S = "abc", shifts = [3,5,9]
	// Output: "rpl"
	// Explanation:
	// We start with "abc".
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
}
