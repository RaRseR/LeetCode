func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        if j, ok := mp[nums[i]]; ok && i - j <= k {
            return true
        }

        mp[nums[i]] = i
    }

    return false
}
