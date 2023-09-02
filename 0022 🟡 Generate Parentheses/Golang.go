func generateParenthesis(n int) (answer []string) {
    helper(&answer, n, 0, 0, "")

    return 
}

func helper(answer *[]string, n int, left int, right int, current string) {
    if left == n && right == n {
        *answer = append(*answer, current)
        return
    }

    if left < n {
        helper(answer, n, left + 1, right, current + "(")
    }

    if left > right {
        helper(answer, n, left, right + 1, current + ")")
    }
}
