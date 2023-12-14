func numSpecial(mat [][]int) (result int) {
    var rows []int = make([]int, len(mat))
    var cols []int = make([]int, len(mat[0]))
    
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            if mat[i][j] == 1 {
                rows[i]++
                cols[j]++
            }
        }
    }

    for i:= 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
                result++
            }
        }
    }

    return
}
