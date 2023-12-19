func imageSmoother(img [][]int) [][]int {
    var result [][]int = make([][]int, len(img))
    for i := 0; i < len(img); i++ {
        result[i] = make([]int, len(img[0]))
    }

    for i := 0; i < len(img); i++ {
        for j := 0; j < len(img[0]); j++ {
            var summary, count int = 0, 0

            for x := i - 1; x <= i + 1; x++ {
                for y := j - 1; y <= j + 1; y++ {
                    if x >= 0 && x < len(img) && y >= 0 && y < len(img[0]) {
                        summary += img[x][y]
                        count++
                    }
                }
            }

            result[i][j] = summary / count
        }
    }

    return result
}
