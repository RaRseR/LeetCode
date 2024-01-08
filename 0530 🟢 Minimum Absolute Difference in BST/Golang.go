func getMinimumDifference(root *TreeNode) int {
    _, _, diff := helper(root)
    return diff
}

func helper(root *TreeNode) (int, int, int) {
    min, max, diff := root.Val, root.Val, math.MaxInt32

    if root.Left != nil {
        lMin, lMax, lDiff := helper(root.Left)
        min = lMin
        if tmpDiff := root.Val - lMax; tmpDiff < diff {
            diff = tmpDiff
        }

        if lDiff < diff {
            diff = lDiff
        }
    }

    if root.Right != nil {
        rMin, rMax, rDiff := helper(root.Right)
        max = rMax
        if tmpDiff := rMin - root.Val; tmpDiff < diff {
            diff = tmpDiff
        }

        if rDiff < diff {
            diff = rDiff
        }
    }

    return min, max, diff
}
