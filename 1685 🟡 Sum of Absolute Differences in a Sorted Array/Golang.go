func getSumAbsoluteDifferences(nums []int) (result []int) {
    var totalSum int = 0
    for i := 0; i < len(nums); i++ {
        totalSum += nums[i]
    }

    var leftSum int = 0

    for i := 0; i < len(nums); i++ {
        var rightSum int = totalSum - leftSum - nums[i]

        var leftCount int = i
        var rightCount int = len(nums) - i - 1

        var leftTotal int = leftCount * nums[i] - leftSum
        var rightTotal int = rightSum - rightCount * nums[i]

        result = append(result, leftTotal + rightTotal)
        leftSum += nums[i]
    }

    return
}
