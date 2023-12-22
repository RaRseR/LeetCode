func maxScore(s string) int {
    var zeros, ones int = 0, 0
    var result int = -1

    for i := 0; i < len(s); i++ {
        if s[i] == 48 {
            zeros++
        } else {
            ones++
        }

        if i != len(s) - 1 {
            result = max(zeros - ones, result)
        }
    }

    return result + ones
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
