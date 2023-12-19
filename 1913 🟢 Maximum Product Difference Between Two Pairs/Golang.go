func maxProductDifference(nums []int) int {
    var w, x, y, z int = math.MinInt, math.MinInt, math.MaxInt, math.MaxInt

    for i := 0; i < len(nums); i++ {
        if nums[i] > w {
            w, x = nums[i], w
        } else if nums[i] > x {
            x = nums[i]
        } else if nums[i] < y {
            y, z = nums[i], y
        } else if nums[i] < z {
            z = nums[i]
        }
    }

    return w * x - y * z
}
