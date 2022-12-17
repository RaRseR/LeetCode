func lengthOfLongestSubstring(s string) (l int) {
    if len(s) == 0 {
        return
    }

    hash := map[byte]int{}

    for i, j := 0, 0; i < len(s); i++ {
        if v, ok := hash[s[i]]; ok && v > j {
            j = v
        }
        if i - j + 1 > l {
            l = i - j + 1
        }
        hash[s[i]] = i + 1
    }

    return
}
