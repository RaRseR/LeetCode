func goodNodes(root *TreeNode) (result int) {
    return dfs(root, root.Val)
}

func dfs(root *TreeNode, max int) int {
    if root == nil {
        return 0
    }

    if root.Val >= max {
        max = root.Val

        return dfs(root.Left, max) + dfs(root.Right, max) + 1
    }

    return dfs(root.Left, max) + dfs(root.Right, max)
}
