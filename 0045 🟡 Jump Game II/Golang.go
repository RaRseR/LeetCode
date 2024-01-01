func jump(nums []int) int {
    var curr int = 0
    var max int = 0
    var result int = 0

    for i := 0; i < len(nums) - 1; i++ {
        if nums[i] + i > max {
            max = nums[i] + i
        }

        if i == curr {
            result++
            curr = max

            if curr >= len(nums) - 1 {
                return result
            }
        }
    }
    
    return 0
}
