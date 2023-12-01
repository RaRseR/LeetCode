func subsets(nums []int) (result [][]int) {
	current := make([]int, 0)

  var backtrack func(idx int)
  backtrack = func(idx int) {
		result = append(result, append([]int{}, current...))

		if idx == len(nums) {
			return
		}

		for i := idx; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i + 1)
			current = current[:len(current)-1]
		}
	}
	backtrack(0)

  return
}
