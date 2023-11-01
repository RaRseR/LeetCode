func findMode(root *TreeNode) (modes []int) {
    var currentValue, currentCount, maxCount int
    
    inOrderTraversal(root, &currentValue, &currentCount, &maxCount, &modes)

    return
}

func inOrderTraversal(node *TreeNode, currentValue, currentCount, maxCount *int, modes *[]int) {
    if node == nil {
        return
    }

    inOrderTraversal(node.Left, currentValue, currentCount, maxCount, modes)

    if node.Val == *currentValue {
        *currentCount++
    } else {
        *currentValue = node.Val
        *currentCount = 1
    }

    if *currentCount == *maxCount {
        *modes = append(*modes, *currentValue)
    } else if *currentCount > *maxCount {
        *maxCount = *currentCount
        *modes = []int{*currentValue}
    }

    inOrderTraversal(node.Right, currentValue, currentCount, maxCount, modes)
}
