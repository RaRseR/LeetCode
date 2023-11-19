func diameterOfBinaryTree(root *TreeNode) (result int) {
    dfs(root, &result)

    return
}

func dfs(root *TreeNode, max *int) int {
    if root == nil {
        return 0
    }

    var maxLeft int = dfs(root.Left, max)
    var maxRight int = dfs(root.Right, max)

    if maxLeft + maxRight > *max {
        *max = maxLeft + maxRight
    }

    if maxLeft > maxRight {
        return maxLeft + 1
    }

    return maxRight + 1
}
