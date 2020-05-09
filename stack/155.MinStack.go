package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt32, stack: []int{}}
}

func (this *MinStack) Push(x int) {
	if this.min >= x {
		this.stack = append(this.stack, this.min)
		this.min = x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	_len := len(this.stack)
	last := this.stack[_len-1]
	this.stack = this.stack[:_len-1]
	if last == this.min {
		this.min = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	fmt.Println("s")
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	print(m.GetMin())
	m.Pop()
	print(m.Top())
	print(m.GetMin())
}
