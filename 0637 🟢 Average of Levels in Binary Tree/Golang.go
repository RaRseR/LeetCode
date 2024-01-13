func averageOfLevels(root *TreeNode) (result []float64) {
    if root == nil {
        return
    }

    queue := []*TreeNode{root}
    
    for len(queue) != 0 {
        currLen := len(queue)
        value := 0.0

        for i := 0; i < currLen; i++ {
            node := queue[0]
            queue = queue[1:]

            if node != nil {
                value += float64(node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, value / float64(currLen))
    }

    return
}
