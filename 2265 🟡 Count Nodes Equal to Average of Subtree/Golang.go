func averageOfSubtree(root *TreeNode) (result int) {
    helper(root, &result)

    return
}

func helper(node *TreeNode, result *int) (int, int) {
    if node == nil {
        return 0, 0
    }
    
    leftSum, leftCount := helper(node.Left, result)
    rightSum, rightCount := helper(node.Right, result)
    
    currSum := node.Val + leftSum + rightSum
    currCount := 1 + leftCount + rightCount
    
    if currSum / currCount == node.Val {
        *result++
    }
    
    return currSum, currCount
}
