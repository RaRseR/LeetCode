func hIndex(citations []int) int {
    var hashMap map[int]int = make(map[int]int)
    var result int = 0

    for i := 0; i < len(citations); i++ {
        hashMap[citations[i]]++
    }

    for i, h := 0, len(citations); i <= h; i++ {
        result = i
        h -= hashMap[i]
    }

    return result
}
