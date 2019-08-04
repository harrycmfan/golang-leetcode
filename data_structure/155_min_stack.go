package data_structure

import "math"

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.


Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
 */


type MinStack struct {
	a []int
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		a: []int{},
		min: math.MaxInt64,
	}
}


func (this *MinStack) Push(x int)  {
	if x <= this.min {
		this.a = append(this.a, this.min)
		this.min = x
	}
	this.a = append(this.a, x)
}


func (this *MinStack) Pop()  {
	pop := this.a[len(this.a)-1]
	if pop == this.min {
		this.min = this.a[len(this.a)-2]
		this.a = this.a[:len(this.a)-2]
	} else {
		this.a = this.a[:len(this.a)-1]
	}
}


func (this *MinStack) Top() int {
	return this.a[len(this.a)-1]
}


func (this *MinStack) GetMin() int {
	return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */