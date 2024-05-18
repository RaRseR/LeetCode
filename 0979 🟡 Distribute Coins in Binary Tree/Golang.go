func distributeCoins(root *TreeNode) (moves int) {
    traverse(root, &moves)

    return
}

func traverse(root *TreeNode, moves *int) int {
    if root == nil {
        return 0
    }

    left := traverse(root.Left, moves)
    right := traverse(root.Right, moves)

    *moves += abs(left) + abs(right)

    return root.Val + left + right - 1
}

func abs(i int) int {
    if i < 0 {
        return - i
    }

    return i
}
