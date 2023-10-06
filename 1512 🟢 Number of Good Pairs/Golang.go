func numIdenticalPairs(nums []int) (result int) {
    pairsMap := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        pairsMap[nums[i]]++

        result += pairsMap[nums[i]] - 1
    }
    
    return
}
