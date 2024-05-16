func largestLocal(grid [][]int) [][]int {
    n := len(grid)

    result := make([][]int, n - 2)

    for row := 0; row < n - 2; row++ {
        result[row] = make([]int, n - 2)

        for col := 0; col < n - 2; col++ {
            max := grid[row][col]
            
            for newRow := row; newRow < row + 3; newRow++ {
                for newCol := col; newCol < col + 3; newCol++ {
                    if grid[newRow][newCol] > max {
                        max = grid[newRow][newCol]
                    }
                }
            }

            result[row][col] = max
        }
    }

    return result
}
