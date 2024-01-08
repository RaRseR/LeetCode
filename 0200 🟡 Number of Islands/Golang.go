func numIslands(grid [][]byte) (result int) {
    for i, row := range grid {
        for j, val := range row {
            if val == '1' {
                helper(&grid, i, j)
                result++
            }
        }
    }
    
    return
}

func helper(grid *[][]byte, i int, j int) {
    if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
        return
    }

    (*grid)[i][j] = '0'

    helper(grid, i - 1, j)
    helper(grid, i + 1, j)
    helper(grid, i, j - 1)
    helper(grid, i, j + 1)
}
