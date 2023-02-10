func removeElement(nums []int, val int) (i int) {
    for j := range nums {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }

    return
}
