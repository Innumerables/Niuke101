package main

import "sort"

//BM55-BM61 递归/回溯

//BM55 没有重复项数字的全排列
var (
	result [][]int
	path   []int
	state  []bool
)

func permute(num []int) [][]int {
	// write code here
	result = make([][]int, 0)
	path = make([]int, 0)
	state = make([]bool, len(num))
	tracking(num, 0)
	return result
}

func tracking(num []int, start int) {
	if len(num) == start {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
	}
	for i := 0; i < len(num); i++ {
		if !state[i] {
			path = append(path, num[i])
			state[i] = true
			tracking(num, start+1)
			state[i] = false
			path = path[:len(path)-1]
		}
	}
}

//BM56 有重复数字的全排列
func permuteUnique(num []int) [][]int {
	result = make([][]int, 0)
	path = make([]int, 0)
	state = make([]bool, len(num))
	sort.Ints(num)
	tracking(num, 0)
	return result
}
func trackings(num []int, start int) {
	if len(num) == start {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
	}
	for i := 0; i < len(num); i++ {
		if i != 0 && num[i] == num[i-1] && !state[i-1] {
			continue
		}
		if !state[i] {
			path = append(path, num[i])
			state[i] = true
			tracking(num, start+1)
			state[i] = false
			path = path[:len(path)-1]
		}
	}
}

var ()

//BM58 字符串的排列
func Permutation(str string) []string {
	// write code here
}
