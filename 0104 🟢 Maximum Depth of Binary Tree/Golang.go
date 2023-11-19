func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    var maxLeft int = maxDepth(root.Left)
    var maxRight int = maxDepth(root.Right)

    if maxLeft > maxRight {
        return maxLeft + 1
    }

    return maxRight + 1
}
