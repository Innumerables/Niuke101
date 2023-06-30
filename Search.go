package main

import (
	"strconv"
	"strings"
)

//牛客101 二分查找/排序，BM17-BM22

// BM17 二分查找 在给定一个升序的序列中找出目标值的位置
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

// BM18 二维数组中的查找
// 利用之前的二分查找，将二维数组取出一维依次的使用二分查找，判断是否存在
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

// BM19 寻找峰值 无序数组峰值元素严格大于左右相邻值的元素，数组可能存在多个峰值，返回任意一个即可
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

// BM20 数组中的逆序对 前面一个数字大于后边的数字就是一个逆序对，求出总数
// 主要思想，利用归并排序中合并的时侯来计算总数，使用全局变量记录总数
var (
	count = 0
)

func InversePairs(nums []int) int {
	// write code here
	l := len(nums)
	if l < 2 {
		return 0
	}
	merge1(nums, 0, l-1)
	return count
}

func merge1(nums []int, left, right int) {
	mid := (left + right) / 2
	if left < right {
		merge1(nums, left, mid)
		merge1(nums, mid+1, right)
		mergeCount(nums, left, mid, right)
	}
}
func mergeCount(nums []int, left, mid, right int) {
	arr := make([]int, right-left+1)
	c := 0
	l := left
	m := mid + 1
	for l <= mid && m <= right {
		if nums[l] <= nums[m] {
			arr[c] = nums[l]
			l += 1
			c += 1
		} else { //当切片前的大于切片右的，因为越靠前越小，因此在此之后到两者交接处都满足>该值，因此加上mid+1-l
			arr[c] = nums[m]
			count = count + mid + 1 - l
			m += 1
			c += 1
		}
	}
	for l <= mid {
		arr[c] = nums[l]
		c += 1
		l += 1
	}
	for m <= right {
		arr[c] = nums[m]
		c += 1
		m += 1
	}
	for k := 0; k < right-left+1; k++ {
		nums[left+k] = arr[k]
	}
}

// BM21 旋转数组中的最小数字，给定一个升序数组，将开始的若干个元素放到末尾，[1,2,3,4,5]变为[3,4,5,1,2]，求这个数组的最小值
// 根据数组的规律，先算中间值，如果中间值大于左右边值，说明最小值存在于mid+1--right之间
// 如果之间值小于最右边值，不能确定该值是否为最小值，right移动到mid,向前测试，两者相等只能减小空间继续
func minNumberInRotateArray(nums []int) int {
	// write code here
	l := len(nums)
	left := 0
	right := l - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

// BM22 比较版本号,遇到0自动跳过
// 主要思想：利用strings.Split 函数将字符串拆分，拆分后返回便是以.分割的切片，然后转换为int类型，比较切片大小即可
func compare(version1 string, version2 string) int {
	// write code here
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	l1 := len(v1)
	l2 := len(v2)
	i := 0
	for i < l1 && i < l2 {
		tmp1, _ := strconv.Atoi(v1[i])
		tmp2, _ := strconv.Atoi(v2[i])
		if tmp1 > tmp2 {
			return 1
		} else if tmp1 < tmp2 {
			return -1
		} else {
			i += 1
		}
	}
	if i == l1 && i == l2 {
		return 0
	} else if i < l1 {
		for i < l1 {
			tmp1, _ := strconv.Atoi(v1[i])
			if tmp1 != 0 {
				return 1
			}
			i += 1
		}
		return 0
	} else {
		for i < l2 {
			tmp2, _ := strconv.Atoi(v2[i])
			if tmp2 != 0 {
				return -1
			}
			i += 1
		}
		return 0
	}
}
