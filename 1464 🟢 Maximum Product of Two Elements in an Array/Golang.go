func maxProduct(nums []int) int {
    var firstNum int
    var secondNum int

    for i := 0; i < len(nums); i++ {
        if nums[i] > firstNum {
            firstNum, secondNum = nums[i], firstNum
        } else if nums[i] > secondNum {
            secondNum = nums[i]
        }
    }

    return (firstNum - 1) * (secondNum - 1)
}
