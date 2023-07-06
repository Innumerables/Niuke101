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

//BM33 二叉树的镜像
func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return nil
	}
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	return pRoot
}

//BM34判断是否是二叉搜索树
func isValidBST(root *TreeNode) bool {
	// write code here
	if root == nil {
		return false
	}
	//引入一个存储二叉树的切片，用来遍历
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil && node.Right != nil {
			if node.Left.Val > node.Val || node.Right.Val < node.Val { //判断此节点的左右子树是否满足
				return false
			}
			if node.Left.Right != nil && node.Left.Right.Val > node.Val { //防止该节点左子树的右节点大于该节点值
				return false
			}
			if node.Right.Left != nil && node.Right.Left.Val < node.Val { //防止该节点的右子树的左节点大于该节点的值
				return false
			}
			queue = append(queue, node.Left, node.Right)
			continue
		} else if node.Left != nil {
			if node.Left.Val > node.Val {
				return false
			}
			queue = append(queue, node.Left)
			continue
		} else if node.Right != nil {
			if node.Right.Val < node.Val {
				return false
			}
			queue = append(queue, node.Right)
		}
	}
	return true
}

//BM35 判断是不是完全二叉树
func isCompleteTree(root *TreeNode) bool {
	// write code here
	queue := []*TreeNode{root}
	flag := 0 //叶子节点终止条件
	//当flag为1时，表明该节点为叶子节点，或为该节点为非满树，此后的节点肯定没有叶子节点。当存在叶子节点时既不是完全二叉树
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right == nil {
			flag = 1
			continue
		} else if node.Left == nil && node.Right != nil {
			return false
		} else if node.Left != nil && node.Right != nil {
			if flag == 1 {
				return false
			}
			queue = append(queue, node.Left, node.Right)
			continue
		} else if node.Left != nil && node.Right == nil {
			if flag == 1 {
				return false
			}
			queue = append(queue, node.Left)
			flag = 1
		}
	}
	return true
}

//第二种解法
func isCompleteTree1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	notComplete := false
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tempQueue := []*TreeNode{}
		for _, node := range queue {
			if node == nil {
				notComplete = true
				continue
			}
			if notComplete {
				return false // 第二次出现空节点
			}
			//nil加紧切片作为判断依据
			tempQueue = append(tempQueue, node.Left)
			tempQueue = append(tempQueue, node.Right)
		}
		queue = tempQueue
	}

	return true
}
