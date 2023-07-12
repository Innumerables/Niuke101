package main

// BM42 用两个栈实现队列
// 栈先进后出，队列先进先出，每次新增一个数时，先将栈中的数据转移到另一个栈中，让后将该数加到该栈中，再将加入后的栈反转即实现了进入队列
// 出队列时只需出栈中的最后一个
var stack1 []int
var stack2 []int

func Push0(node int) {
	if len(stack1) == 0 {
		stack1 = append(stack1, node)
		return
	} else {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
		stack2 = append(stack2, node)
		for i := len(stack2) - 1; i >= 0; i-- {
			stack1 = append(stack1, stack2[i])
		}
		stack2 = []int{}
		return
	}
}

func Pop0() int {
	l := len(stack1)
	num := stack1[l-1]
	stack1 = stack1[:l-1]
	return num
}

// BM43 包含min函数的栈
func Push(node int) {
	stack1 = append(stack1, node)
}
func Pop() {
	if len(stack1) == 0 {
		return
	}
	stack1 = stack1[:len(stack1)-1]
	return
}
func Top() int {
	return stack1[len(stack1)-1]
}
func Min() int {
	if len(stack1) == 1 {
		return stack1[0]
	}
	min := stack1[len(stack1)-1]
	stack2 = append(stack2, min)
	for i := len(stack1) - 2; i >= 0; i-- {
		if stack1[i] < min {
			min = stack1[i]
		}
		stack2 = append(stack2, stack1[i])
	}
	stack1 = []int{}
	for i := len(stack2) - 1; i >= 0; i-- {
		stack1 = append(stack1, stack2[i])
	}
	stack2 = []int{}
	return min
}

// BM44 有序括号序列
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	mid := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	res := []byte{}
	for i := 0; i < len(s); i++ {
		v, ok := mid[s[i]]
		if ok {
			res = append(res, v)
		} else {
			if len(res) == 0 || s[i] != res[len(res)-1] {
				return false
			}
			res = res[:len(res)-1]
		}
	}
	if len(res) != 0 {
		return false
	}
	return true
}
