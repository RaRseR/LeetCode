func getMaximumGold(grid [][]int) int {
    result := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            result = max(result, findMaxGold(grid, i, j))
        }
    }

    return result
}

func findMaxGold(grid [][]int, i, j int) int {
    if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == 0 {
        return 0
    }

    directions := [5]int{0, 1, 0, -1, 0}

    originalGold := grid[i][j] 
    grid[i][j] = 0

    maxGold := 0
    for k := 0; k < 4; k++ {
        maxGold = max(maxGold, findMaxGold(grid, directions[k] + i, directions[k + 1] + j))
    }

    grid[i][j] = originalGold
    return maxGold + originalGold
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
