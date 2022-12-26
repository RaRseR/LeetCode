func maxSubArray(nums []int) (sum int) {
    var local int

    for _, num := range nums {
        local += num

        if sum < local {
          sum = local
        }

        if local < 0 {
            local = 0
        }
      }
    return 
}
