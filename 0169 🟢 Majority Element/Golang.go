func majorityElement(nums []int) int {
    if len(nums) <= 2 {
        return nums[0]
    }

    var appearance map[int]int = make(map[int]int)

    for i := 0; i < len(nums); i++ {
        appearance[nums[i]]++

        if appearance[nums[i]] > len(nums) / 2 {
            return nums[i]
        }
    }

    return 0
}
