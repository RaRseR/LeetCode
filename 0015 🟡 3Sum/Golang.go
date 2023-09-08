func threeSum(nums []int) (result [][]int) {
    sort.Ints(nums)

    for idx1 := 0; idx1 < len(nums) - 2; idx1++ {
        if idx1 > 0 && nums[idx1] == nums[idx1 - 1] {
            continue
        }

        idx2 := idx1 + 1
        idx3 := len(nums) - 1

        for idx2 < idx3 {
            sum := nums[idx1] + nums[idx2] + nums[idx3]

            if sum == 0 {
                result = append(result, []int{nums[idx1], nums[idx2], nums[idx3]})

                idx3--

                for idx2 < idx3 && nums[idx3] == nums[idx3 + 1] {
                    idx3--
                }
            } else if sum > 0 {
                idx3--
            } else {
                idx2++
            }
        }
    }

    return
}
