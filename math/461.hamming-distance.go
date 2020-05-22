/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	res := 0
	a := x^y
	for a!=0{
		if a%2==1{
			res+=1
		}
		a = a>>1
	}
	return res
}
// @lc code=end

