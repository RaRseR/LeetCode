func findMin(nums []int) int {
    if len(nums) == 0 {
        return nums[0]
    }

    left := 0
    right := len(nums) - 1

    for left < right {
        mid := (left + right) / 2

        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return nums[left]
}
