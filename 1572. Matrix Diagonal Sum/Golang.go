func diagonalSum(mat [][]int) int {
    l := len(mat)
    
    res := 0
    
    for i := 0; i < l; i++ {
        res += mat[i][i] + mat[l - 1 - i][i]
    }
    
    if (l % 2 == 0) {
        return res
    }
    
    return res - mat[l/2][l/2];
}
