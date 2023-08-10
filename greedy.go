package main

import "sort"

//BM95-BM96

//BM95分糖果问题
func candy(arr []int) int {
	// write code here
	n := len(arr)
	if n < 2 {
		return 1
	}
	left := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = 1
	}
	for i := 0; i < n-1; i++ {
		if arr[i] < arr[i+1] {
			left[i+1] = left[i] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			left[i-1] = max(left[i-1], left[i]+1)
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += left[i]
	}
	return sum
}

//BM96主持人调度
func minmumNumberOfHost(n int, startEnd [][]int) int {
	// write code here
	start := make([]int, n)
	end := make([]int, n)
	for i := 0; i < n; i++ {
		start[i] = startEnd[i][0]
		end[1] = startEnd[i][1]
	}
	sort.Ints(start)
	sort.Ints(end)
	sum, j := 0, 0
	for i := 0; i < n; i++ {
		if start[i] >= end[j] {
			j++
		} else {
			sum++
		}
	}
	return sum
}
