package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	d := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		d[i] = make([]int, l2+1)
	}
	return dfs(word1, word2, l1, l2, d)
}

func dfs(word1, word2 string, l1, l2 int, d [][]int) int {
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}
	if d[l1][l2] > 0 {
		return d[l1][l2]
	}
	ans := 0
	if word1[l1-1] == word2[l2-1] {
		ans = dfs(word1, word2, l1-1, l2-1, d)
	} else {
		ans = min(dfs(word1, word2, l1-1, l2-1, d), min(dfs(word1, word2, l1-1, l2, d), dfs(word1, word2, l1, l2-1, d))) + 1
	}
	d[l1][l2] = ans
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {

	// Input: word1 = "horse", word2 = "ros"
	// Output: 3
	// Explanation:
	// horse -> rorse (replace 'h' with 'r')
	// rorse -> rose (remove 'r')
	// rose -> ros (remove 'e')
	// Input: word1 = "intention", word2 = "execution"
	// Output: 5
	// Explanation:
	// intention -> inention (remove 't')
	// inention -> enention (replace 'i' with 'e')
	// enention -> exention (replace 'n' with 'x')
	// exention -> exection (replace 'n' with 'c')
	// exection -> execution (insert 'u')
	fmt.Println(minDistance("horse", "ros"))
}
