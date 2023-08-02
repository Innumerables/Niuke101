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
		if i != 0 && num[i] == num[i-1] && !state[i-1] { //去重操作
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

//BM57 岛屿数量
func solvess(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				//在当前的岛屿连接所有岛屿，记为1
				dfs(grid, i, j, m, n)
			}
		}
	}
	return res
}
func dfs(grid [][]byte, i, j, m, n int) {
	if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0' //将所有相关岛屿都变为0，以防再次计数
	dfs(grid, i-1, j, m, n)
	dfs(grid, i+1, j, m, n)
	dfs(grid, i, j-1, m, n)
	dfs(grid, i, j+1, m, n)
}

var (
	restr   []string
	pathstr []byte
	states  []bool
)

//BM58 字符串的排列
func Permutation(str string) []string {
	// write code here
	restr = make([]string, 0)
	states = make([]bool, len(str))
	midstr := []byte(str)
	sort.Slice(midstr, func(i, j int) bool {
		return midstr[i] < midstr[j]
	})
	dfstr(string(midstr), 0)
	return restr
}
func dfstr(str string, start int) {
	if len(str) == start {
		tmp := string(pathstr)
		restr = append(restr, tmp)
	}
	for i := 0; i < len(str); i++ {
		if i != 0 && str[i] == str[i-1] && !states[i-1] {
			continue
		}
		if !states[i] {
			pathstr = append(pathstr, str[i])
			states[i] = true
			dfstr(str, start+1)
			states[i] = false
			pathstr = pathstr[:len(pathstr)-1]
		}
	}
}

//BM59 N皇后问题
var ans int

func Nqueen(n int) int {
	track := make([][]string, n)
	for i := 0; i < n; i++ {
		track[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			track[i][j] = "."
		}
	}
	backtracking(n, 0, track)
	return ans
}
func backtracking(n, row int, track [][]string) {
	if n == row {
		ans++
		return
	}
	for col := 0; col < n; col++ { //以列为开始，当行走到最后时说明满足
		if !isRight(n, row, col, track) {
			continue
		}
		track[row][col] = "Q"
		backtracking(n, row+1, track)
		track[row][col] = "."
	}
}
func isRight(n, row, col int, track [][]string) bool { //判断是否合法
	for i := 0; i < row; i++ { //判断当前列是否含有Q
		if track[i][col] == "Q" {
			return false
		}
	}
	//判断四十五度角是否含有Q
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	return true
}

//BM60 括号生成
func generateParenthesis(n int) []string {
	// write code here
	res := []string{}
	var dfs func(l, r int, path string)
	dfs = func(l, r int, path string) {
		if len(path) == 2*n {
			res = append(res, path)
			return
		}
		if l > 0 {
			dfs(l-1, r, path+"(")
		}
		if r > l { //只有当右括号大于左括号时加入右括号才能达到合法括号，先左后右
			dfs(l, r-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}

//BM61 矩阵最长递增路径
func solve111(matrix [][]int) int {
	// write code here
	m := len(matrix)
	n := len(matrix[0])
	used := make([][]int, m)
	maxpath := 1
	for i := 0; i < m; i++ {
		used[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxV := dfsss(matrix, i, j, used)
			maxpath = max(maxpath, maxV)
		}
	}
	return maxpath
}

func dfsss(matrix [][]int, i, j int, used [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	if used[i][j] != 0 {
		return used[i][j]
	}
	maxV := 1
	if i-1 > 0 && matrix[i-1][j] > matrix[i][j] {
		maxV = max(maxV, dfsss(matrix, i-1, j, used)+1)
	}
	if i+1 < m && matrix[i+1][j] > matrix[i][j] {
		maxV = max(maxV, dfsss(matrix, i+1, j, used)+1)
	}
	if j-1 > 0 && matrix[i][j-1] > matrix[i][j] {
		maxV = max(maxV, dfsss(matrix, i, j-1, used)+1)
	}
	if j+1 < n && matrix[i][j+1] > matrix[i][j] {
		maxV = max(maxV, dfsss(matrix, i, j+1, used)+1)
	}
	used[i][j] = maxV
	return maxV
}
