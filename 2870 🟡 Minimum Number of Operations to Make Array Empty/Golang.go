func minOperations(nums []int) int {
    var occurance map[int]int = make(map[int]int)

    for _, value := range nums {
        occurance[value]++
    }

    var result int = 0
    for _, value := range occurance {
        if value == 1 {
            return -1
        }

        result += int(math.Ceil(float64(value) / 3))
    }

    return result
}
