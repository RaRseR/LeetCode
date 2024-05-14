func matrixScore(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    result := (1 << (n - 1)) * m

    for i := 1; i < n; i++ {
        current := 0

        for j := 0; j < m; j++ {
            if grid[j][i] == grid[j][0] {
                current++ 
            }
        }

        result += max(current, m - current) * ( 1 << (n - i - 1))
    }

    return result
}


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
