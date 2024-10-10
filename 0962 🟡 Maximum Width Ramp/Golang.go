func maxWidthRamp(nums []int) int {
    n := len(nums)
    monotonicStack := make([]int, 0, n)

    for i := 0; i < n; i++ {
        if len(monotonicStack) == 0 || nums[monotonicStack[len(monotonicStack) - 1]] > nums[i] {
            monotonicStack = append(monotonicStack, i)
        }
    }

    maxWidth := 0

    for i := n - 1; i >= 0; i-- {
        for len(monotonicStack) > 0 && nums[monotonicStack[len(monotonicStack) - 1]] <= nums[i] {
            j := monotonicStack[len(monotonicStack) - 1]

            monotonicStack = monotonicStack[:len(monotonicStack) - 1]

            if i - j > maxWidth {
                maxWidth = i - j
            }
        }
    }

    return maxWidth
}
