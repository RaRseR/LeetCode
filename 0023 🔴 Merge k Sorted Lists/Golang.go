func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 || (len(lists) == 1 && lists[0] == nil) {
        return nil
    }

    for len(lists) > 1 {
        mergedLists := []*ListNode{}

        for i := 0; i < len(lists); i += 2 {
            list1 := lists[i]
            var list2 *ListNode

            if i + 1 < len(lists) {
                list2 = lists[i + 1]
            }

            mergedLists = append(mergedLists, mergeLists(list1, list2))
        }

        lists = mergedLists
    }

    return lists[0]
}

func mergeLists(list1, list2 *ListNode) *ListNode {
    dummy := &ListNode{
        Val: 0,
        Next: nil,
    } 

    tail := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }

        tail = tail.Next
    }

    if list1 != nil {
        tail.Next = list1
    }
    if list2 != nil {
        tail.Next = list2
    }

    return dummy.Next
}
