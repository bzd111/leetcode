package main

import "math/rand"

func rand10() int {
	index := 100
	for index > 39 {
		index = 7*(rand7()-1) + rand7() - 1
	}
	return index%10 + 1
}

func rand7() int {
	return rand.Intn(7) + 1
}
