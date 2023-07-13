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

//BM45 滑动窗口最大值
func maxInWindows(num []int, size int) []int {
	// write code here
	mid := []int{}
	res := []int{}

	for i := 0; i < len(num); i++ {
		for len(mid) > 0 && mid[len(mid)-1] < num[i] {
			mid = mid[:len(mid)-1]
		}
		mid = append(mid, num[i])
		if i >= size && mid[0] == num[i-size] { //当该最大值在滑动窗口外时，删除该元素
			mid = mid[1:]
		}
		if i >= size-1 {
			res = append(res, mid[0])
		}
	}
	return res
}

//BM46 最小的k个数
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	input = MergeSort(input) //利用归并排序得到排好序的切片
	return input[:k]
}

//BM 47 寻找第k大的数
func findKth(a []int, n int, K int) int {
	// write code here
	a = MergeSort(a) //利用归并排序得到排好序的切片
	return a[K-1]
}
