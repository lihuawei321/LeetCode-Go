package main

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

//给定这个链表：1->2->3->4->5
//当 k = 2 时，应当返回: 2->1->4->3->5
//当 k = 3 时，应当返回: 3->2->1->4->5

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 1 {
		return head
	}
	var pre *ListNode
	cur := head
	var next *ListNode
	check := head
	canProceed := 0
	count := 0
	//检查长度是否满足翻转
	for canProceed < k && check != nil {
		check = check.Next
		canProceed++
	}
	//满足条件，进行翻转
	if canProceed == k {
		for count < k && cur != nil {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			count++
		}
		if next != nil {
			//head为翻转后的尾结点
			head.Next = reverseKGroup(next, k)
		}
		//pre为翻转后的头结点
		return pre
	} else {
		return head
	}
}
func main() {

}
