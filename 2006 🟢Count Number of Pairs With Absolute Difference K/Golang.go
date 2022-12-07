func countKDifference(nums []int, k int) (c int) {
    m := map[int]int{}

    for _, n := range nums {
        m[n]++
        c += m[n + k] + m[n - k]
    }
    
    return c
}
