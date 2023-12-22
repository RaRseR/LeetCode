func rotate(nums []int, k int) {
    k %= len(nums)

    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(nums []int) {
    var start, end int = 0, len(nums) - 1

    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
