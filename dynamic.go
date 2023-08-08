package main

import (
	"math"
	"strconv"
	"strings"
)

//牛客网101动态规划BM62-BM82题解

// BM62斐波那契数列
func Fibonacci(n int) int {
	// write code here
	if n <= 2 {
		return 1
	}
	a, b, c := 1, 1, 1
	for i := 2; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

// BM63 跳台阶
var sum int

func jumpFloor(number int) int {
	// write code here
	sum = 0
	jumpFloor(number)
	return sum
}
func jump(n int) {
	if n == 0 {
		sum++
		return
	}
	jump(n - 1)
	if n > 1 {
		jump(n - 2)
	}
}

func jumpFloor1(number int) int { //动态规划解法
	dp := make([]int, number+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}

// BM64 最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	// write code here
	length := len(cost)
	var dp [1000000]int
	for i := 2; i <= length; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[length]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// BM65 最长公共子序列
func LCS1(s1 string, s2 string) string {
	// write code here
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	var res string
	for i, j := m, n; dp[i][j] > 0; {
		if s1[i-1] == s2[j-1] {
			res = string(s1[i-1]) + res
			i--
			j--
		} else if dp[i-1][j] >= dp[i][j-1] {
			i--
		} else if dp[i-1][j] < dp[i][j-1] {
			j--
		}
	}
	if dp[m][n] == 0 {
		return "-1"
	}
	return res
}

// BM66最长公共子串
func LCS(str1 string, str2 string) string {
	// write code here
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	res := 0
	end := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if res < dp[i][j] {
				end = i
				res = dp[i][j]
			}
		}
	}
	if res == 0 {
		return ""
	}
	return str1[end-res : end]
}

// BM67不同路径的数目
func uniquePaths(m int, n int) int {
	// write code here
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// BM68 矩阵最小路径和
func minPathSum(matrix [][]int) int {
	// write code here
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[m-1][n-1]
}

// BM69 把数字翻译成字符串
func solverrr(nums string) int {
	// write code here
}

// BM70 兑换零钱
// dp[i] 定义组成数字i最小的硬币数量，
func minMoney(arr []int, aim int) int {
	// write code here
	dp := make([]int, aim+1)
	dp[0] = 0
	for i := 1; i <= aim; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(arr); i++ {
		for j := arr[i]; j <= aim; j++ {
			if dp[j-arr[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-arr[i]]+1) //递推公式,j的组成方式多种多样，选取最小的硬币数量
			}
		}
	}
	if dp[aim] == math.MaxInt32 {
		return -1
	}
	return dp[aim]
}

// BM71 最长上升子序列
func LIS(arr []int) int {
	// write code here
	m := len(arr)
	if m == 0 {
		return 0
	}
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	maxS := 1
	for i := 1; i < m; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxS {
			maxS = dp[i]
		}
	}
	return maxS
}

// BM72连续子数组的最大和
func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	m := len(array)
	dp := make([]int, m)
	dp[0] = array[0]
	mx := array[0]
	for i := 1; i < m; i++ {
		dp[i] = max(array[i], dp[i-1]+array[i])
		mx = max(mx, dp[i])
	}
	return mx
}

// BM73最长回文字串
func getLongestPalindrome(A string) int {
	// write code here
	m := len(A)
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
	}
	max := 0
	for i := m - 1; i >= 0; i-- {
		for j := i; j < m; j++ {
			if A[i] == A[j] {
				if j-i > max {
					max = j - i
				}
				dp[i][j] = true
			} else if dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	return max + 1
}

// BM74 数字字符串转换为IP地址
func restoreIpAddresses(s string) []string {
	// write code here
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	res := []string{}
	var dfs func(start int, path []string)
	dfs = func(start int, path []string) {
		if len(path) == 4 && start == n {
			res = append(res, strings.Join(path, "."))
			return
		}
		if start == n {
			return
		}
		if s[start] == '0' {
			dfs(start+1, append(path, "0"))
			return
		}
		for i := start; i < n && i <= start+2; i++ {
			s := s[start : i+1]
			if isValids(s) {
				dfs(i+1, append(path, s))
			}
		}
	}
	dfs(0, []string{})
	return res
}
func isValids(str string) bool {

	sum, _ := strconv.Atoi(str)

	return sum <= 255 && sum > 0
}

// BM75编辑距离
func editDistance(str1 string, str2 string) int {
	// write code here
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Mins(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
func Mins(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
}

// BM78打家劫舍1
func rob(nums []int) int {
	// write code here
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// BM79打家劫舍2
// 划分不同区间执行两次取最大值
func rob2(nums []int) int {
	a := rob(nums[0 : len(nums)-1])
	b := rob(nums[1:])
	return max(a, b)

}

// BM80 买卖股票的最好时机,一次买卖
func maxProfit1(prices []int) int {
	// write code here
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

// BM81 买卖股票的最好时机，多次买卖
func maxProfit2(prices []int) int {
	// write code here
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

// BM82 买卖股票的最好时机，最多两次买卖
func maxProfit(prices []int) int {
	// write code here
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}
	return dp[len(prices)-1][3]
}

// k次买卖
func maxProfits(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}
