func isBalanced(root *TreeNode) bool {
    result, _ := dfs(root)
    
    return result
}

func dfs(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }

    isLeftBalanced, leftHeight := dfs(root.Left)
    isRightBalanced, rightHeight := dfs(root.Right)

    if isLeftBalanced && isRightBalanced && abs(leftHeight - rightHeight) <= 1{
        if leftHeight > rightHeight {
            return true, leftHeight + 1
        }

        return true, rightHeight + 1
    }

    return false, -1
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
