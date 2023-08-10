package main

//BM97-BM101

//BM97旋转数组
func solve12(n int, m int, a []int) []int {
	// write code here
	if m == 0 {
		return a
	}
	move := m % n
	res := a[n-move:]
	res = append(res, a[:n-move]...)
	return res
}
