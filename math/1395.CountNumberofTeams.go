package main

import "fmt"

func numTeams(rating []int) int {
	ans := 0
	for j := 0; j < len(rating); j++ {
		l := 0
		r := 0
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				l++
			}
		}
		for k := j + 1; k < len(rating); k++ {
			if rating[j] < rating[k] {
				r++
			}
		}
		ans += l*r + (j-l)*(len(rating)-1-r-j)
	}
	return ans
}

func main() {
	rating := []int{2, 5, 3, 4, 1} // 3
	fmt.Println((numTeams(rating)))
	rating = []int{2, 1, 3} // 0
	fmt.Println((numTeams(rating)))
	rating = []int{1, 2, 3, 4} // 4
	fmt.Println((numTeams(rating)))
}
