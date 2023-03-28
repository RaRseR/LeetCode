func strStr(haystack string, needle string) int {
	n := len(needle)

	for i := 0; i <= len(haystack) - n; i++ {
		j := 0

		for j < n {
			if haystack[i + j] != needle[j] {
				break
			}
			j++
		}

		if j == n {
			return i
		}
	}

	return -1
}
