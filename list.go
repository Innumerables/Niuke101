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
