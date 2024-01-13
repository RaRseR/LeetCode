func zigzagLevelOrder(root *TreeNode) (result [][]int) {
    if root == nil {
        return
    }

    queue := []*TreeNode{root}
    direction := true

    for len(queue) != 0 {
        currLen := len(queue)
        zigzag := make([]int, 0, currLen)

        for i := 0; i < currLen; i++ {
            node := queue[0]
            queue = queue[1:]

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }

            if direction {
                zigzag = append(zigzag, node.Val)
            } else {
                zigzag = append([]int{node.Val}, zigzag...)
            }
        }

        direction = !direction

        result = append(result, zigzag)
    }

    return
}
