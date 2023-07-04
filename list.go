package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// BM1 反转链表
func ReverseList(pHead *ListNode) *ListNode {
	var dhead *ListNode
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	// midlistNode := &ListNode{}
	for pHead.Next != nil {
		midlistNode := pHead.Next //保存下一个节点，以便反向指向
		pHead.Next = dhead
		dhead = pHead
		pHead = midlistNode
	}
	pHead.Next = dhead
	return pHead
}

// BM2 链表内指定区间反转,找到前一个在区间内反转
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	thead := &ListNode{}
	thead.Next = head
	pre := thead
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	mhead := pre.Next
	mthead := mhead
	var dhead *ListNode
	for i := m; i <= n; i++ {
		mid := mhead.Next
		mhead.Next = dhead
		dhead = mhead
		mhead = mid
	}
	pre.Next = dhead
	mthead.Next = mhead
	return thead.Next
}

// BM2 采用头插法的方式将后方的节点依次插入到指定位置M的后方
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	thead := &ListNode{}
	thead.Next = head
	pre := thead
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	mhead := pre.Next
	for i := m; i < n; i++ {
		tmp := mhead.Next
		mhead.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}
	return thead.Next
}

// BM3  链表中的节点每k个一组反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	thead := &ListNode{}
	thead.Next = head
	i := 1
	for head != nil {
		head = head.Next
		i++
	}
	if head == nil || i < k {
		return head
	}
	head = thead.Next
	mhead := thead
	for j := 0; j < i/k; j++ { //判断循环几次，当剩余节点数小于k时，不变
		for t := 1; t < k; t++ { //利用头插法逆转对应k的顺序
			tmp := head.Next
			head.Next = tmp.Next
			tmp.Next = mhead.Next
			mhead.Next = tmp
		}
		//将节点移至到下一个k前一个节点
		mhead = head
		head = head.Next
	}
	return thead.Next
}

// BM4 合并两个排序的链表，两个链表为递增的
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	thead := &ListNode{}
	zhead := thead
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			thead.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			thead.Next = pHead2
			pHead2 = pHead2.Next
		}
		thead = thead.Next
	}
	if pHead1 != nil {
		thead.Next = pHead1
	} else {
		thead.Next = pHead2
	}
	return zhead.Next
}

// BM5 合并k个升序的链表,利用上述两个链表合并，依次合并k个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	k := len(lists)
	var head *ListNode
	if k == 0 {
		return head
	}
	head = lists[0]
	for i := 1; i < k; i++ {
		head = Merge(head, lists[i])
	}
	return head
}

// BM6 判断链表中是否有环
func hasCycle(head *ListNode) bool {
	// write code here
	if head == nil {
		return false
	}
	mhead := head
	khead := head
	for khead != nil && khead.Next != nil {
		mhead = mhead.Next
		khead = khead.Next.Next
		if mhead == khead {
			return true
		}
	}
	return false
}

// BM7 判断出链表中环的入口点
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead.Next == nil {
		return nil
	}
	mhead := pHead
	khead := pHead
	flag := 0
	for khead != nil && khead.Next != nil {
		mhead = mhead.Next
		khead = khead.Next.Next
		if mhead == khead {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return nil
	}
	mhead = pHead
	for mhead != khead {
		mhead = mhead.Next
		khead = khead.Next
	}
	return mhead
}

// BM8 链表中倒数最后k个节点
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	thead := pHead
	i := 1
	for pHead != nil {
		pHead = pHead.Next
		i += 1
	}
	if k > i {
		return thead
	}
	pHead = thead
	for j := 1; j+k != i; j++ {
		pHead = pHead.Next
	}
	return pHead.Next
}

// BM8 链表中倒数第k个节点，双指针法
func FindKthToTail2(pHead *ListNode, k int) *ListNode {
	if pHead == nil {
		return pHead

	}
	thead := pHead
	for i := 0; i < k; i++ {
		if pHead == nil {
			return pHead
		}
		pHead = pHead.Next
	}
	for pHead != nil {
		thead = thead.Next
		pHead = pHead.Next
	}
	return thead
}

// BM9 删除链表中的倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	if head == nil {
		return nil
	}
	thead := &ListNode{}
	thead.Next = head
	t := 0
	for head != nil {
		head = head.Next
		t += 1
	}
	if t < n {
		return thead.Next
	}
	mid := thead
	for j := 0; j+n != t; j++ {
		mid = mid.Next
	}
	mid.Next = mid.Next.Next
	return thead.Next
}

// BM10 两个链表中的第一个公共的节点
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	p1, p2 := 0, 0
	mid1 := pHead1
	mid2 := pHead2
	for mid1 != nil {
		p1 += 1
		mid1 = mid1.Next
	}
	for mid2 != nil {
		p2 += 1
		mid2 = mid2.Next
	}
	if p1 > p2 {
		t := p1 - p2
		for t > 0 && pHead1 != nil {
			pHead1 = pHead1.Next
			t -= 1
		}
		for pHead1 != pHead2 {
			pHead1 = pHead1.Next
			pHead2 = pHead2.Next
		}
	} else {
		t := p2 - p1
		for t > 0 && pHead2 != nil {
			pHead2 = pHead2.Next
			t -= 1
		}
		for pHead1 != pHead2 {
			pHead1 = pHead1.Next
			pHead2 = pHead2.Next
		}
	}
	return pHead1
}

// BM11 两个链表相加，利用之前的反转链表，将两个链表反转，然后相加计算值
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	thead := &ListNode{}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	head1 = ReverseList(head1) //链表反转函数ReverseList
	head2 = ReverseList(head2)
	midsum := 0
	midhead := thead
	for head1 != nil || head2 != nil {
		midnode := &ListNode{Val: midsum}
		if head1 != nil {
			midnode.Val = midnode.Val + head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			midnode.Val = midnode.Val + head2.Val
			head2 = head2.Next
		}
		//以此来获得是否大于10,是否向前进位
		midsum = midnode.Val / 10
		midnode.Val = midnode.Val % 10 //获得小于10的当前节点的数值
		midhead.Next = midnode
		midhead = midhead.Next
	}
	if midsum > 0 { //判断最前位是否有进位
		midnode := &ListNode{Val: midsum}
		midhead.Next = midnode
	}
	thead = ReverseList(thead.Next)
	return thead
}

//BM12 单链表的排序

func sortInList1(head *ListNode) *ListNode { //超时算法，选择排序O（n^2)
	// write code here
	thead := &ListNode{}
	thead.Next = head
	for head.Next != nil {
		fast := head.Next
		min := head
		for fast != nil {
			if fast.Val < min.Val {
				min = fast
			}
			fast = fast.Next
		}
		tmp := min.Val
		min.Val = head.Val
		head.Val = tmp
		head = head.Next
	}
	return thead.Next
}

// BM12 单链表的排序 分治法归并排序利用数组进行分治法归并排序所使用的MergeSort在sort.go中实现
func sortInList(head *ListNode) *ListNode {
	n := 0
	midhead := head
	for head != nil {
		head = head.Next
		n += 1
	}
	if n == 1 {
		return midhead
	}
	midval := make([]int, n)
	i := 0
	for midhead != nil {
		midval[i] = midhead.Val
		midhead = midhead.Next
		i = i + 1
	}
	midval = MergeSort(midval)
	thead := &ListNode{
		Val:  midval[0],
		Next: nil,
	}
	midhead = thead
	for i := 1; i < len(midval); i++ {
		tmp := &ListNode{
			Val:  midval[i],
			Next: nil,
		}
		thead.Next = tmp
		thead = thead.Next
	}
	return midhead
}

// BM12 单链表排序，直接在链表上使用分治法归并排序
func sortInList2(head *ListNode) *ListNode {
	// write code here
	if head == nil {
		return head
	}
	return splitLink(head)
}
func splitLink(head *ListNode) *ListNode {
	// 拆分为只剩下一个节点或者空节点
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	slow, fast := head, head
	// 通过快慢指针把链表拆分为2部分
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// pre为slow的前一个，将next指向为空断开
	pre.Next = nil
	// 使用递归方式分别拆分2部分链表
	left := splitLink(head)
	right := splitLink(slow)
	return mergeLink(left, right)
}
func mergeLink(l1, l2 *ListNode) *ListNode {
	// 合并两个有序的链表
	// 创建新的空头节点
	newH := &ListNode{}
	node := newH
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	// 直接拼接还没有遍历完，剩下的链表
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return newH.Next
}

// BM13 判断链表是否为回文结构,借助数组，将链表值按顺序存入数组中，判断数组中值是否为回文结构
func isPail(head *ListNode) bool {
	// write code here
	if head == nil {
		return true
	}
	n := 0
	midhead := head
	for head != nil { //在此可以直接将节点值通过midaval = append(midval,head.val)进行赋值，但其可能消耗的内存空间较大
		n += 1 //因为当切片当内村不足时，<1024时会以当前的两倍进行扩增，也会进行数据迁移复制，因此不如得到大小进行再赋值
		head = head.Next
	}
	if n == 1 {
		return true
	}
	midval := make([]int, n)
	i := 0
	for midhead != nil {
		midval[i] = midhead.Val
		i += 1
		midhead = midhead.Next
	}
	for i := 0; i < n/2; i++ {
		if midval[i] != midval[n-i-1] {
			return false
		}
	}
	return true
}

// BM14 链表的奇偶重排，意思是奇偶排在一起，先奇数后偶数
func oddEvenList(head *ListNode) *ListNode { //利用两次循环
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	thead := &ListNode{}
	n := 1
	mhead := thead
	ohead := head
	for head != nil {
		if n%2 != 0 {
			tmp := &ListNode{Val: head.Val}
			thead.Next = tmp
			thead = thead.Next
		}
		n += 1
		head = head.Next
	}
	n = 1
	for ohead != nil {
		if n%2 == 0 {
			tmp := &ListNode{Val: ohead.Val}
			thead.Next = tmp
			thead = thead.Next
		}
		n += 1
		ohead = ohead.Next
	}
	return mhead.Next
}
func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddHead := &ListNode{}
	evenHead := &ListNode{}
	odd := oddHead
	even := evenHead
	n := 1
	for head != nil { //利用一次循环
		tmp := &ListNode{Val: head.Val}
		if n%2 != 0 {
			oddHead.Next = tmp
			oddHead = oddHead.Next
		} else {
			evenHead.Next = tmp
			evenHead = evenHead.Next
		}
		n += 1
		head = head.Next
	}
	oddHead.Next = even.Next
	return odd.Next
}

// BM15 删除有序列表中的重复元素
func deleteDuplicatesI(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	thead := head
	pre := head.Next
	for head != nil && head.Next != nil {
		for head.Val == pre.Val {
			head.Next = pre.Next
			pre = pre.Next
			if pre == nil {
				break
			}
		}
		head = head.Next
	}
	return thead
}

// BM16 删除链表中所有重复出现的元素，只保留链表中只出现一次的元素
func deleteDuplicates(head *ListNode) *ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	//preH 为当前值的前一个，以便移出去相同的值；mid为当前值的后一个，判断是否相等，直到找到不同的值
	preH := &ListNode{}
	preH.Next = head
	thead := preH
	mid := head.Next
	for head != nil && head.Next != nil {
		flag := 0
		for head.Val == mid.Val {
			flag = 1
			mid = mid.Next
			if mid == nil { //当到链表末尾时，意味着从当前到末尾全部一样，即舍去，将前一个指针指向空返回链表
				preH.Next = nil
				return thead.Next
			}
		}
		if flag == 1 { //判断是否有相等的，若有，则直接将preH 移到该不同值的前一个，head指向不同的值位置
			preH.Next = mid
			head = mid
		} else { //没有相等的则每一个向后移动一步
			preH = preH.Next
			head = head.Next
		}
		mid = head.Next
	}
	return thead.Next
}
