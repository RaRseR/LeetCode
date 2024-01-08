func summaryRanges(nums []int) (result []string) {
    if len(nums) == 0 {
        return
    }

    start := 0

    for i := 0; i < len(nums); i++ {
        if i < len(nums) - 1 && nums[i] + 1 == nums[i + 1] {
            continue
        }

        if start == i {
            result = append(result, fmt.Sprintf("%d", nums[i]))
        } else {
            result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i]))
        }

        start = i + 1
    }

    return
}
