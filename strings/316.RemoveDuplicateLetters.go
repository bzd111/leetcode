package main

import "fmt"

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	used := [26]bool{}
	res := []rune{}
	for i := range s {
		count[s[i]-'a']++
	}

	for i := range s {
		count[s[i]-'a']--
		if used[s[i]-'a'] {
			continue
		}
		for len(res) > 0 && res[len(res)-1] > rune(s[i]) && count[res[len(res)-1]-'a'] > 0 {
			used[res[len(res)-1]-'a'] = false
			res = res[:len(res)-1]
		}
		res = append(res, rune(s[i]))
		used[s[i]-'a'] = true
	}
	return string(res)
}

func main() {

	// Input: "bcabc"
	// Output: "abc"

	// Input: "cbacdcbc"
	// Output: "acdb"

	// Input: "cbbbcaa"
	// Output: "bca"

	// Input: "bbcaac""
	// Output: "bac"
	s := "bcabc"
	fmt.Println(removeDuplicateLetters(s))

	s = "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s))

	s = "cbbbcaa"
	fmt.Println(removeDuplicateLetters(s))

	s = "bbcaac"
	fmt.Println(removeDuplicateLetters(s))
}
