func reorderList(head *ListNode)  {
    p1, p2 := head, head.Next

    for p2 != nil && p2.Next != nil {
        p1 = p1.Next
        p2 = p2.Next.Next
    }

    p2 = p1.Next
    p1.Next = nil

    var previous *ListNode

    for p2 != nil {
        tmp := p2.Next
        p2.Next = previous
        previous = p2
        p2 = tmp
    }

    p1, p2 = head, previous

    for p2 != nil {
        tmp1, tmp2 := p1.Next, p2.Next

        p1.Next = p2
        p2.Next = tmp1

        p1, p2 = tmp1, tmp2
    }
}
