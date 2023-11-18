func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    hashMap := map[*Node]*Node{}

    node := head
    for node != nil {
        hashMap[node] = &Node{
            Val: node.Val,
            Next: nil,
            Random: nil,
        }

        node = node.Next
    }

    node = head
    for node != nil {
        hashMap[node].Next = hashMap[node.Next]
        hashMap[node].Random = hashMap[node.Random]

        node = node.Next
    }

    return hashMap[head]
}
