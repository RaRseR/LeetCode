func flatten(root *TreeNode) {
	var helper func(root *TreeNode)

	var node *TreeNode
	helper = func (root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Right)
		helper(root.Left)
		root.Right = node
		root.Left = nil
		node = root
	}
	helper(root)
}
