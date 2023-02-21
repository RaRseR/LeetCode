func searchInsert(nums []int, target int) int {		
	l := 0
	r := len(nums) - 1

	for l < r {
	  m := l + (r - l) / 2
			
		if nums[m] == target {
      return m
    } else if nums[m] > target {
      r = m
    } else {
      l = m + 1
    }
  }
	
  if nums[l] < target {
    return l + 1
  }
  return l
}
