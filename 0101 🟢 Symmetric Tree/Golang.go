func isSymmetric(root *TreeNode) bool {
    return helper(root.Left, root.Right)
}

func helper(leftNode, rightNode *TreeNode) bool {
    if leftNode == nil && rightNode == nil {
        return true
    }

    if (leftNode == nil || rightNode == nil) || leftNode.Val != rightNode.Val {
        return false
    }

    return helper(leftNode.Left, rightNode.Right) && helper(leftNode.Right, rightNode.Left)
}
