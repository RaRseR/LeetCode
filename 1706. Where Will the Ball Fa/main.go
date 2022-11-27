func findBall(grid [][]int) []int {
    rows, cols := len(grid), len(grid[0])
    result := make([]int, cols)
    
    for col := 0; col < cols; col++ {
        currentCol := col
        
        for currentRow := 0; currentRow < rows; currentRow++ {
            
            nextCol := currentCol + grid[currentRow][currentCol]
            
            if nextCol < 0 || nextCol >= cols || grid[currentRow][currentCol] != grid[currentRow][nextCol] {
                currentCol = -1
                break;
            }
            
            currentCol = nextCol
        }
        
        result[col] = currentCol;
    }
    return result
}
