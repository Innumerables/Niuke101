package main

//二叉树相关题目
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BM23 二叉树的前序遍历\
//递归的方式实现
func preorderTraversal(root *TreeNode) []int {
	// write code here
	res := make([]int, 0)
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) //前中后序遍历取决于该行放的位置，在此时前序遍历
		preOrder(node.Left)
		// res = append(res, node.Val) //中序遍历
		preOrder(node.Right)
		// res = append(res, node.Val) //后序遍历
	}
	preOrder(root)
	return res
}

// func preOrder(res *[]int, root *TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	*res = append(*res, root.Val)
// 	preOrder(res, root.Left)
// 	preOrder(res, root.Right)
// }

//非递归的方式实现，前序遍历
func preorderTraversalNo(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	zn := make([]*TreeNode, 0)
	res := make([]int, 0)
	zn = append(zn, root)
	for len(zn) != 0 {
		res = append(res, zn[len(zn)-1].Val)
		no := zn[len(zn)-1]
		zn = zn[:len(zn)-1]
		if no.Right != nil {
			zn = append(zn, no.Right)
		}
		if no.Left != nil {
			zn = append(zn, no.Left)
		}
	}
	return res
}

//BM26 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	arry := []*TreeNode{}
	arry = append(arry, root)
	for len(arry) > 0 {
		midarry := []*TreeNode{} //用于记录下一层的叶子
		midres := []int{}
		for _, node := range arry {
			midres = append(midres, node.Val)
			if node.Left != nil {
				midarry = append(midarry, node.Left)
			}
			if node.Right != nil {
				midarry = append(midarry, node.Right)
			}
		}
		res = append(res, midres)
		arry = midarry //移动到下一层
	}
	return res
}

//BM27 按之字形顺序打印二叉树
func Print(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return [][]int{}
	}
	res := [][]int{}
	arry := []*TreeNode{}
	arry = append(arry, pRoot)
	i := 1
	for len(arry) > 0 {
		midarry := []*TreeNode{} //用于记录下一层的叶子
		midres := []int{}
		for _, node := range arry {
			midres = append(midres, node.Val)
			if node.Left != nil {
				midarry = append(midarry, node.Left)
			}
			if node.Right != nil {
				midarry = append(midarry, node.Right)
			}
		}
		if i%2 == 0 { //当为偶数层时反转所记录的叶子节点的值以实现之字形记录
			for j := 0; j < len(midres)/2; j++ {
				midres[j], midres[len(midres)-1-j] = midres[len(midres)-1-j], midres[j]
			}
		}
		res = append(res, midres)
		arry = midarry //移动到下一层
		i += 1
	}
	return res
}

//BM28 二叉树最大的深度
func maxDepth(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}
	res := 0
	arry := []*TreeNode{}
	arry = append(arry, root)
	for len(arry) > 0 {
		midarry := []*TreeNode{} //用于记录下一层的节点
		for _, node := range arry {
			if node.Left != nil {
				midarry = append(midarry, node.Left)
			}
			if node.Right != nil {
				midarry = append(midarry, node.Right)
			}
		}
		res += 1       //用于记录层数
		arry = midarry //移动到下一层
	}
	return res
}

//递归的方式实现获得最大深度
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	r := max(left, right) + 1
	return r
}
func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

//BM29 二叉树中和为某一值的路径
func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
