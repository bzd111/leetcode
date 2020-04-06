package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	length := 0
	start := 0
	for i := range s {
		cur := int(math.Max(float64(getLens(s, i, i)), float64(getLens(s, i, i+1))))
		if cur > length {
			length = cur
			start = i - (cur-1)/2
		}

	}
	return s[start : start+length]
}

func getLens(s string, l int, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l -= 1
		r += 1
	}
	// r - l -2 + 1
	return r - l - 1
}

func main() {
	t := "babad"
	fmt.Println(longestPalindrome(t))
	t = "cbbd"
	fmt.Println(longestPalindrome(t))
	t = "a"
	fmt.Println(longestPalindrome(t))
	t = "ccc"
	fmt.Println(longestPalindrome(t))
}
