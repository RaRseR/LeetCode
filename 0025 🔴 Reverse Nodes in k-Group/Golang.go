func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	count := 0

	for count < k {
		if node == nil {
			return head
		}
		node = node.Next
		count++
	}

	prev := reverseKGroup(node, k)

	for count > 0 {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
		count--
	}

	return prev
}
