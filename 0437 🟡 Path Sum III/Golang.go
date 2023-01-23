func pathSum(root *TreeNode, targetSum int) (totalPaths int) {
    if root == nil {
        return
    }
    
    helper(root, make([]int, 0), targetSum, &totalPaths)
    
    return
}

func helper(node *TreeNode, paths []int, target int, totalPaths *int) []int {
    for i := range paths {
        paths[i] += node.Val
    }
    
    paths = append(paths, node.Val)
    
    if node.Left != nil {
        paths = helper(node.Left, paths, target, totalPaths)
    }
    if node.Right != nil {
        paths = helper(node.Right, paths, target, totalPaths)
    }
    
    for i := range paths {
        if paths[i] == target {
            *totalPaths += 1
        }
        paths[i] -= node.Val
    }
    paths = paths[:len(paths)-1]
    
    return paths
}
