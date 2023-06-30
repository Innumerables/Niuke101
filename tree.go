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
