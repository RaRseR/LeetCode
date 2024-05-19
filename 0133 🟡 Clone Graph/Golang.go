func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }

    clones := map[int]*Node{}

    return helper(node, clones)
}

func helper(node *Node, clones map[int]*Node) *Node {
    if clone, ok := clones[node.Val]; ok {
        return clone
    }

    nodeClone := &Node{Val: node.Val}
    clones[node.Val] = nodeClone

    for _, neighbor := range node.Neighbors {
        nodeClone.Neighbors = append(nodeClone.Neighbors, helper(neighbor, clones))
    }

    return nodeClone
}
