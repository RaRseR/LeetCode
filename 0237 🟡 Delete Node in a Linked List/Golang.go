func deleteNode(node *ListNode) {
    if node != nil {
      *node = *node.Next
    }

    return
}
