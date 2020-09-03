package main

import "strconv"

/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start
func calPoints(ops []string) int {
	result := []int{}
	for _, op := range ops {
		switch op {
		case "C":
			result = result[:len(result)-1]
		case "D":
			result = append(result, result[len(result)-1]*2)

		case "+":
			result = append(result, result[len(result)-1]+result[len(result)-2])
		default:
			num, _ := strconv.Atoi(op)
			result = append(result, num)
		}
	}
	return sum(result)
}

func sum(list []int) int {
	s := 0
	for _, v := range list {
		s += v
	}
	return s
}

// @lc code=end
