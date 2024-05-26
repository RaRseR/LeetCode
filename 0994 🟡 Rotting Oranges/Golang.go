func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    queue := [][2]int{}
    fresh := 0

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 1 {
                fresh++
            }
            if grid[r][c] == 2 {
                queue = append(queue, [2]int{r, c})
            }
        }
    }

    directions := [][2]int{{1,0}, {-1, 0}, {0, 1}, {0, -1}}
    result := 0

    for len(queue) > 0 && fresh > 0 {
        for _, orange := range queue {
            queue = queue[1:]
            for _, dir := range directions {
                r, c := orange[0] + dir[0], orange[1] + dir[1]

                if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != 1 {
                    continue
                }

                grid[r][c] = 2
                queue = append(queue, [2]int{r, c})

                fresh--
            }
        }

        result++
    }

    if fresh == 0 {
        return result
    }

    return -1
}
