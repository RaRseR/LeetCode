func arrayPairSum(nums []int) (result int) {
    sort.Ints(nums)

    for i := 0; i < len(nums); i += 2{
        result += nums[i]
    }

    return result
}
