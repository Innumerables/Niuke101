package main

import "sort"

//双指针算法，BM87-BM94

//BM87 合并两个有序数组
func merges(A []int, m int, B []int, n int) {
	// write code here
	if m == 0 { //当A为空时，将B中的数据转到A中
		for i := 0; i < n; i++ {
			A[m] = B[i]
			m++
		}
		return
	}
	j := 0
	for i := 0; i < n; i++ {
		if A[j] > B[i] { //当A中数据大于B时，将A中数据后移，将B插入
			for a := m - 1; a >= j; a-- {
				A[a+1] = A[a]
			}
			A[j] = B[i]
			m += 1
		} else if j == m-1 { //当A中比较的当前数为A中最后一个数时，说明B中剩下的数都比A大，将其存入后边即可
			for b := i; b < n; b++ {
				j++
				A[j] = B[b]
			}
			return
		} else { //当A中比较数小于B中数据时，A中位置前移，B位置不变
			i--
			j++
		}
	}
}

//BM88 判断是否为回文字符串
func judge(str string) bool {
	// write code here
	ll := len(str)
	b := str[:]
	for i := 0; i < ll/2; i++ {
		if b[i] != b[ll-1-i] { //首尾对称比较
			return false
		}
	}
	return true
}

//BM89 合并区间
type Interval struct {
	Start int
	End   int
}

func merge0(intervals []*Interval) []*Interval {
	// write code here
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	var res []*Interval
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= res[len(res)-1].End {
			res[len(res)-1].End = max(intervals[i].End, res[len(res)-1].End)
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

//BM90最小覆盖字串
func minWindow(S string, T string) string {
	// write code here
	s := S[:]
	t := T[:]
	tmap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tmap[t[i]]++

	}
	smap := map[byte]int{}
	mid := []int{}
	flag := 0
	minlen := len(s)
	start := 0
	end := 0
	tr := 0
	for i := 0; i < len(s); i++ {
		_, ok := tmap[s[i]]
		if ok {
			smap[s[i]]++
			mid = append(mid, i) //用于存储与T值相对应所有S中的坐标，便与向前跳跃匹配
			if smap[s[i]] == tmap[s[i]] {
				flag++
			}
		}
		for flag == len(tmap) {
			tr = 1
			if mid[len(mid)-1]-mid[0] < minlen { //获得最小的距离
				minlen = mid[len(mid)-1] - mid[0]
				start = mid[0]
				end = mid[len(mid)-1]
			}
			smap[s[mid[0]]]--                      //减去满足条件时第一个位置所存的字符的次数
			if smap[s[mid[0]]] < tmap[s[mid[0]]] { //若减去之后次数小于T中次数，则说明不满足条件跳出循环
				flag--
			}
			mid = mid[1:] //然后删除第一个位置，若减去次数满足则继续循环，直到不满足，找到最小的
		}
	}
	if tr == 0 {
		return ""
	}
	return string(s[start : end+1])
}

//BM91 反转字符串
func solve11(str string) string {
	// write code here
	s := []byte(str)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return string(s)
}

//BM92 最长无重复子数组
func maxLength(arr []int) int {
	// write code here
	mid := map[int]int{}
	max := 0
	t := 0
	for i := 0; i < len(arr); i++ {
		v, ok := mid[arr[i]]
		if ok {
			t = max1(v+1, t) //若相等则判断，该相等的位置和最开始的位置谁大，取大的计算当前长度
		}
		max = max1(max, i-t+1) //计算最长长度
		mid[arr[i]] = i
	}
	return max
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//BM93 盛水最多的容器
func maxArea(height []int) int {
	// write code here
	left := 0
	right := len(height)
	max := 0
	for left < right {
		min := minx(height[left], height[right])
		sum := min * (right - left)
		if sum > max {
			max = sum
		}
		if height[left] < height[right] { //左边小移动左边
			left++
		} else {
			right--
		}
	}
	return max
}

func minx(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//BM94 接雨水问题
