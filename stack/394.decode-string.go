package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	var numStack []int
	var strStack []string
	tail := ""
	for i := 0; i < len(s); i++ {
		switch {
		case isDigit(string(s[i])):
			n, _ := strconv.Atoi(string(s[i]))
			for isDigit(string(s[i+1])) {
				n = 10*n + int(s[i+1])
				i++
			}
			numStack = append(numStack, n)
		case s[i] == '[':
			strStack = append(strStack, tail)
			tail = ""
		case s[i] == ']':
			tmp := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			times := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			tmp += strings.Repeat(tail, times)
			tail = tmp
		default:
			tail += string(s[i])
		}
	}
	fmt.Println(strStack)
	fmt.Println(numStack)
	return tail
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}
	return true
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

// @lc code=end
