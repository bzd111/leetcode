package main

import (
	"fmt"
	"reflect"
)

/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

// @lc code=start
func backspaceCompare(S string, T string) bool {
	s := realString(S)
	t := realString(T)
	return reflect.DeepEqual(s, t)
}

func realString(s string) []rune {
	result := []rune{}
	for _, i := range s {
		if i == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab#c", "ad##c"))
}
