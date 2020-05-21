package main

import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	pairs := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := []string{}
	for _, i := range s {
		d := string(i)
		if _, ok := pairs[d]; !ok {
			stack = append(stack, d)
			continue
		}

		if len(stack) > 0 && stack[len(stack)-1] == pairs[d] {
			stack = stack[0 : len(stack)-1]
			continue
		}
		stack = append(stack, d)

	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// @lc code=end

func main() {
	s := "[]"
	fmt.Println(isValid(s))
}
