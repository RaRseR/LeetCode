func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }

    var leftString string = ""
    var rightString string = ""

    if root.Left != nil || root.Right != nil {
        leftString = "(" + tree2str(root.Left) + ")"
    }

    if root.Right != nil {
        rightString = "(" + tree2str(root.Right) + ")"
    }

    return fmt.Sprintf("%v", root.Val) + leftString + rightString
}
