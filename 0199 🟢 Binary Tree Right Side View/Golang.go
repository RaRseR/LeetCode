func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    var result []int = make([]int, 0)

    var queue []*TreeNode = make([]*TreeNode, 1)
    queue[0] = root

    for len(queue) != 0 {
        var currentLength int = len(queue)
        var rightSide *TreeNode

        for i := 0; i < currentLength; i++ {
            var node *TreeNode = queue[0]
            queue = queue[1:]

            if node != nil {
                rightSide = node

                queue = append(queue, node.Left)
                queue = append(queue, node.Right)
            }
        }

        if rightSide != nil {
            result = append(result, rightSide.Val)
        }
    }

    return result
}
