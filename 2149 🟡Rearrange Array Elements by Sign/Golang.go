func rearrangeArray(nums []int) []int {
    n := len(nums)
    answer := make([]int, n)
    positiveIndex, negativeIndex := 0, 1

    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            answer[positiveIndex] = nums[i]
            positiveIndex += 2
        } else {
            answer[negativeIndex] = nums[i]
            negativeIndex += 2
        }
    }

    return answer
}
