func levelOrder(root *TreeNode) (result [][]int) {
    if root == nil {
        return
    }

    var queue []*TreeNode = make([]*TreeNode, 1)
    var level int = 0
    queue[0] = root

    for len(queue) != 0 {      
        var currentLength int = len(queue)

        for i := 0; i < currentLength; i++ {
            var node *TreeNode = queue[0]
            queue = queue[1:]

            if len(result) <= level {
                result = append(result, []int{node.Val})
            } else {
                result[level] = append(result[level], node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        level++
    }

    return 
}
