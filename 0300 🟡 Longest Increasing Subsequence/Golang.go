func lengthOfLIS(nums []int) int {
    var result []int = make([]int, 0, len(nums) / 2)

    for i := 0; i < len(nums); i++ {
        if len(result) == 0 || nums[i] > result[len(result) - 1] {
            result = append(result, nums[i])
            continue
        }

        var left, right, mid int = 0, len(result), 0
        for left != right {
            mid = left + (right - left) / 2

            if nums[i] > result[mid] {
                left = mid + 1
            } else {
                right = mid
            }
        }

        result[right] = nums[i]
    }

    return len(result)
}
