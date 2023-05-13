func inorderTraversal(root *TreeNode) (ans []int) {
    var helper func(*TreeNode) 
    
    helper = func(root *TreeNode) {
        if root == nil {return}
        
        helper(root.Left)

        ans = append(ans, root.Val)

        helper(root.Right)
        
    }

    helper(root)
    
    return 
}
