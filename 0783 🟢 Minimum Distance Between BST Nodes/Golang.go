func minDiffInBST(root *TreeNode) int {
    min, prev := math.MaxInt32, -1

    dfs(root, &min, &prev)

    return min
}

func dfs(node *TreeNode, min *int, prev *int) {
    if node == nil {
        return
    }

    dfs(node.Left, min, prev)

    if *prev != -1 && node.Val - *prev < *min {
        *min = node.Val - *prev
    }

    *prev = node.Val

    dfs(node.Right, min, prev)
}
