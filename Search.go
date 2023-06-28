package main

//牛客101 二分查找/排序，BM17-BM22

//BM17 二分查找 在给定一个升序的序列中找出目标值的位置
func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	left := 0
	right := l - 1
	if nums[left] > target {
		return -1
	}
	if nums[right] < target {
		return -1
	}
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
			continue
		} else {
			right = mid - 1
			continue
		}
	}
	return -1
}

//BM18 二维数组中的查找
//利用之前的二分查找，将二维数组取出一维依次的使用二分查找，判断是否存在
func Find(target int, array [][]int) bool {
	l := len(array)
	for i := 0; i < l; i++ {
		result := search(array[i], target)
		if result != -1 {
			return true
		}
	}
	return false
}

//BM19 寻找峰值 无序数组峰值元素严格大于左右相邻值的元素，数组可能存在多个峰值，返回任意一个即可
func findPeakElement(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] < nums[i+1] {
			if i+1 == l-1 || nums[i+1] > nums[i+2] {
				return i + 1
			}
		}
	}
	return 0
}

//BM20 数组中的逆序对 前面一个数字大于后边的数字就是一个逆序对，求出总数
func InversePairs(nums []int) int {
	// write code here
}
