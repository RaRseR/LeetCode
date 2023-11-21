func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return areEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func areEqual(root, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil || root.Val != subRoot.Val {
		return false
	}

	return areEqual(root.Left, subRoot.Left) && areEqual(root.Right, subRoot.Right)
}
