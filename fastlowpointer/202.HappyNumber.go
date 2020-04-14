package main

import "fmt"

func isHappy(n int) bool {
	slow := n
	fast := n
	slow = calc(slow)
	fast = calc(fast)
	fast = calc(fast)
	for slow != fast {
		slow = calc(slow)
		fast = calc(fast)
		fast = calc(fast)
	}
	if slow == 1 {
		return true
	}
	return false
}

func calc(n int) int {
	sum := 0
	for n != 0 {
		m := n % 10
		sum += m * m
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}
