package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
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

// 链表内指定区间反转,找到前一个在区间内反转
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

// 采用头插法的方式将后方的节点依次插入到指定位置M的后方
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

// 链表中的节点每k个一组反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	thead := &ListNode{}
	thead.Next = head
	i := 1
	for head != nil {
		head = head.Next
		i++
	}
	head = thead.Next
	mhead := thead
	for j := 0; j < i/k; j++ {
		for t := 0; t < k; t++ {
			tmp:=
		}
	}
}
