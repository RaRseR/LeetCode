func rangeSumBST(root *TreeNode, low int, high int) int {
    var helper func(root *TreeNode) int
    helper = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        sum := helper(root.Left) + helper(root.Right)

        if root.Val >= low && root.Val <= high {
            return sum + root.Val
        }

        return sum
    }

    return helper(root)
}
