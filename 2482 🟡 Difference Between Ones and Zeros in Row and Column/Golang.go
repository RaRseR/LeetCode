func onesMinusZeros(grid [][]int) [][]int {
    var rows []int = make([]int, len(grid))
    var cols []int = make([]int, len(grid[0]))

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                rows[i]++
                cols[j]++
            } else {
                rows[i]--
                cols[j]--
            }
        }
    }

    var result [][]int = make([][]int, len(grid))

    for i := 0; i < len(grid); i++ {
        result[i] = make([]int, len(grid[0])) 
        for j := 0; j < len(grid[i]); j++ {
            result[i][j] = rows[i] + cols[j]
        }
    }

    return result
}
