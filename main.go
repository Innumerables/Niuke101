package main

import (
	"fmt"
)

// 生成相应的二维数组
func createArry(l, m int) [][]int {
	arry := [][]int{}
	for i := 0; i < l; i++ {
		tmp := []int{}
		for j := 0; j < m; j++ {
			tmp = append(tmp, j)
		}
		arry = append(arry, tmp)
	}
	return arry
}

// 牛客101
func test_CompletePack1(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = maxs(dp[j], dp[j-weight[i]]+value[i])
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

// test_CompletePack2 先遍历背包, 在遍历物品
func test_CompletePack2(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	// j从0 开始
	for j := 0; j <= bagWeight; j++ {
		for i := 0; i < len(weight); i++ {
			if j >= weight[i] {
				// 推导公式
				dp[j] = maxs(dp[j], dp[j-weight[i]]+value[i])
			}
			// debug
			//fmt.Println(dp)
		}
	}
	return dp[bagWeight]
}

func maxs(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{1, 1, 1}
	price := []int{1, 1, 1}
	fmt.Println(test_CompletePack1(weight, price, 4))
	fmt.Println(test_CompletePack2(weight, price, 4))
}
