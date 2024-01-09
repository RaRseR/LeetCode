func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    sequence1 := make([]int, 0)
    sequence2 := make([]int, 0)

    dfs(root1, &sequence1)
    dfs(root2, &sequence2)

    if len(sequence1) != len(sequence2) {
        return false
    }

    for i := 0; i < len(sequence1); i++ {
        if sequence1[i] != sequence2[i] {
            return false
        }
    }

    return true
}

func dfs(root *TreeNode, sequence *[]int) {
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        *sequence = append(*sequence, root.Val)
    }

    dfs(root.Left, sequence)
    dfs(root.Right, sequence)
}
