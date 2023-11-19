func reductionOperations(nums []int) (result int) {
    slices.Sort(nums)

    for i := len(nums) - 1; i > 0; i-- {
        if nums[i] > nums[i - 1] {
            result += len(nums) - i
        }
    }
    
    return
}
