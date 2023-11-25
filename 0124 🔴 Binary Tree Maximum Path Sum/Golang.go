func maxPathSum(root *TreeNode) int {
	var result int = math.MinInt32

	helper(root, &result)

	return result
}

func helper(root *TreeNode, currentMax *int) int {
    if root == nil {
        return 0
    }

	left := max(0, helper(root.Left, currentMax))
	right := max(0, helper(root.Right, currentMax))
	*currentMax = max(*currentMax, root.Val+left+right)

	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
