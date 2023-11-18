func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carryOver := 0

    result := &ListNode{
        Val: 0,
        Next: nil,
    }

    for pointer := result; l1 != nil || l2 != nil || carryOver > 0; pointer = pointer.Next {
        if l1 != nil {
            carryOver += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            carryOver += l2.Val
            l2 = l2.Next
        }

        pointer.Next = &ListNode{
            Val: carryOver%10,
            Next: nil,
        }
        carryOver /= 10
    }

    return result.Next
}
