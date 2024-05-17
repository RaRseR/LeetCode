func partition(s string) (result [][]string) {
    partition := make([]string, 0)

    helper(s, &result, &partition)

    return
}

func helper(s string, result *[][]string, partition *[]string) {
    if len(s) == 0 {
        newPartition := make([]string, len(*partition))
        copy(newPartition, *partition)

        *result = append(*result, newPartition)
        return
    }

    for i := 1; i <= len(s); i++ {
        prefix := s[:i]
        postfix := s[i:]

        if isPalindrome(prefix) {
            *partition = append(*partition, prefix)

            helper(postfix, result, partition)

            *partition = (*partition)[:len(*partition) - 1]
        }
    }

    return
}

func isPalindrome(subString string) bool {
    left, right := 0, len(subString) - 1

    for left < right {
        if subString[left] != subString[right] {
            return false
        }

        left++
        right--
    }

    return true
}
