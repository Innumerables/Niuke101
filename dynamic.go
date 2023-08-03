package main

//牛客网101动态规划BM62-BM82题解

//BM62斐波那契数列
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

//BM63 跳台阶
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

//BM64 最小花费爬楼梯
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

//BM67不同路径的数目
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

//BM68 矩阵最小路径和
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
