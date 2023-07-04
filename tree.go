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
	left := maxDepth1(root.Left)
	right := maxDepth1(root.Right)
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
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val) //减去所经过的节点的值
}

//BM30 二叉搜索树于双向链表
//利用分治法来合并二叉搜索树为双向链表
func Convert(pRootOfTree *TreeNode) *TreeNode {
	// write code here
	if pRootOfTree == nil {
		return nil
	}
	left := splitTree(pRootOfTree)
	//找到最左边的节点
	for left.Left != nil {
		left = left.Left
	}
	return left
}

//分树阶段直到叶子节再向上合并
func splitTree(pRootOfTree *TreeNode) *TreeNode {
	//判断是否为叶子节点
	if pRootOfTree.Left == nil && pRootOfTree.Right == nil {
		return pRootOfTree
	}
	left := pRootOfTree.Left
	right := pRootOfTree.Right
	//判断是否为空，为空则不进行搜索
	if left != nil {
		left = splitTree(left)
	}
	if right != nil {
		right = splitTree(right)
	}
	return mergeTree(pRootOfTree, left, right)
}

//合并阶段
func mergeTree(pRootOfTree, left, right *TreeNode) *TreeNode {
	//判断left是否为空，为空则不需要找到最右边的节点连接
	if left != nil {
		for left.Right != nil {
			left = left.Right
		}
	}
	if left != nil {
		left.Right = pRootOfTree
		pRootOfTree.Left = left
	}
	if right != nil {
		pRootOfTree.Right = right
		right.Left = pRootOfTree
	}
	//每次合并返回的都是左叶子节点，在下次合并时，需要找到左子树的最后一个节点来继续合并当前的树
	if pRootOfTree.Left != nil {
		return pRootOfTree.Left
	} else {
		return pRootOfTree
	}
}

//BM31 对称的二叉树
func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return false
	}
	return isSym(pRoot, pRoot)
}

func isSym(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right != nil) {
		return false
	}
	return left.Val == right.Val && isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
}

//BM32 合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// write code here
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
