package main

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
