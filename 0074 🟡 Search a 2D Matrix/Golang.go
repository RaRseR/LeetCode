func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    left, right := 0, m * n - 1

    for left <= right {
        middle := (left + right) / 2
        val := matrix[middle / n][middle % n]

        if val == target {
            return true
        } else if val < target {
            left = middle + 1
        } else {
            right = middle - 1
        }
    }
    return false
}
