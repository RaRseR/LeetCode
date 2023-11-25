func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    var mid int = indexOf(preorder[0], inorder)

    return &TreeNode{
        Val: preorder[0],
        Left: buildTree(preorder[1:mid + 1], inorder[:mid]),
        Right: buildTree(preorder[mid + 1:], inorder[mid + 1:]),
    }
}

func indexOf(num int, nums []int) int {
    for i, v := range nums {
        if v == num {
            return i
        }
    }

    return -1
}
