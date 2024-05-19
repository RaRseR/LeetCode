func maxAreaOfIsland(grid [][]int) (result int) {
    current := 0

    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if grid[r][c] == 1 {
                current = helper(&grid, r, c)

                if current > result {
                    result = current
                }
            }
        }
    }

    return
}

func helper(grid *[][]int, r, c int) int {
    if r < 0 || r >= len(*grid) || c < 0 || c >= len((*grid)[0]) || (*grid)[r][c] == 0 {
        return 0
    }

    (*grid)[r][c] = 0

    return helper(grid, r + 1, c) +
        helper(grid, r - 1, c) +
        helper(grid, r, c + 1) +
        helper(grid, r, c - 1) + 
        1
}
