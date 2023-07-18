package main

import "sort"

//BM50-54

//BM50 两数之和
func twoSum(numbers []int, target int) []int {
	// write code here
	mid := map[int]int{}
	res := []int{}
	for i := 0; i < len(numbers); i++ {
		v, ok := mid[target-numbers[i]]
		if !ok {
			mid[numbers[i]] = i
			continue
		}
		v = +1
		i += 1
		res = append(res, v, i)
	}
	return res
}

//BM51 数组中出现次数超过一半的数字
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	mid := map[int]int{}
	for i := 0; i < len(numbers); i++ {
		_, ok := mid[numbers[i]]
		if ok {
			mid[numbers[i]] += 1
		} else {
			mid[numbers[i]] = 1
		}
	}

	res := -1
	var result int
	for i, v := range mid {
		if v > res {
			res = v
			result = i
		}
	}
	if res > len(numbers)/2 {
		return result
	}
	return 0
}

//BM52 数组中只出现一次的两个数字
func FindNumsAppearOnce(nums []int) []int {
	// write code here
	mid := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := mid[nums[i]]
		if ok {
			mid[nums[i]] += 1
		} else {
			mid[nums[i]] = 1
		}
	}

	res := 1
	result := []int{}
	for i, v := range mid {
		if v == res {
			res = v
			result = append(result, i)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

//BM53 缺失的第一个整数
func minNumberDisappeared(nums []int) int {
	// write code here
	mid := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := mid[nums[i]]
		if ok {
			mid[nums[i]] += 1
		} else {
			mid[nums[i]] = 1
		}
	}
	i := 0
	for i = 1; i <= len(nums); i++ {
		_, ok := mid[i]
		if !ok {
			return i
		}
	}
	return i
}

//BM54 三数之和
func threeSum(num []int) [][]int {
	// write code here
	sort.Ints(num)
	res := [][]int{}
	// 找出a + b + c = 0
	// a = nums[i], b = nums[left], c = nums[right]
	for i := 0; i < len(num)-2; i++ {
		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		n1 := num[i]
		if n1 > 0 {
			break
		}
		// 去重a
		if i > 0 && n1 == num[i-1] {
			continue
		}
		l, r := i+1, len(num)-1
		for l < r {
			n2, n3 := num[l], num[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
				for l < r && num[l] == n2 {
					l++
				}
				for l < r && num[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
