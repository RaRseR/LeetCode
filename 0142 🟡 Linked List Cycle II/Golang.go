func detectCycle(head *ListNode) *ListNode {
    var fast *ListNode = head
    var slow *ListNode = head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            break
        }
    }

    if fast == nil || fast.Next == nil {
        return nil
    }
    
    fast = head

    for fast != slow {
        fast = fast.Next
        slow = slow.Next 
    }

    return fast
}
