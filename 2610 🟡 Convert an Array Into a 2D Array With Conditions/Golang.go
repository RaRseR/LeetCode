func findMatrix(nums []int) (result [][]int) {
    var frequency map[int]int = make(map[int]int)

    for _, value := range nums {
        if frequency[value] >= len(result) {
            result = append(result, []int{})
        }

        result[frequency[value]] = append(result[frequency[value]], value)
        frequency[value]++
    }

    return
}
