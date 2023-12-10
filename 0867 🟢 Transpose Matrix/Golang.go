func transpose(matrix [][]int) [][]int {
    var m, n int = len(matrix), len(matrix[0])

    var result [][]int = make([][]int, n)
    for i := range result {
        result[i] = make([]int, m)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            result[j][i] = matrix[i][j]
        }
    }

    return result
}
