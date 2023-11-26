func getSumAbsoluteDifferences(nums []int) (result []int) {
    var totalSum int = 0
    for i := 0; i < len(nums); i++ {
        totalSum += nums[i]
    }

    var leftSum int = 0

    for i := 0; i < len(nums); i++ {
        var rightSum int = totalSum - leftSum - nums[i]

        var leftTotal int = i * nums[i] - leftSum
        var rightTotal int = rightSum - (len(nums) - i - 1) * nums[i]

        result = append(result, leftTotal + rightTotal)
        leftSum += nums[i]
    }

    return
}
