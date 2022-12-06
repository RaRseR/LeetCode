func subsets(nums []int) (result [][]int) {
    sort.Ints(nums)

	for i := 0; i <= len(nums); i++ {
		helper(nums, i, 0, []int{}, &result)
	}

	return result
}

func helper(n []int, i int, s int, p []int, r *[][]int) {
	if len(p) == i {
		*r = append(*r, append([]int{}, p...))
		return
	}

	for j := s; j <= len(n) - (i - len(p)); j++ {
		p = append(p, n[j])
		helper(n, i, j + 1, p, r)
		p = p[:len(p) - 1]
	}
}
