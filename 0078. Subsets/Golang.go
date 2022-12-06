func subsets(nums []int) (result [][]int) {
    sort.Ints(nums)

	for i := 0; i <= len(nums); i++ {
		helper(nums, i, 0, []int{}, &result)
	}

	return result
}

func helper(n []int, k int, s int, p []int, r *[][]int) {
	if len(p) == k {
		*r = append(*r, append([]int{}, p...))
		return
	}

	for i := s; i <= len(n) - (k - len(p)); i++ {
		p = append(p, n[i])
		helper(n, k, i + 1, p, r)
		p = p[:len(p) - 1]
	}
}
