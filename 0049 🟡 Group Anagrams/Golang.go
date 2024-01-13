func groupAnagrams(strs []string) [][]string {
    mp := make(map[[26]int][]string)

    for i := 0; i < len(strs); i++ {
        key := [26]int{}

        for j := 0; j < len(strs[i]); j++ {
            key[strs[i][j] - 'a']++
        }

        mp[key] = append(mp[key], strs[i])
    }

    result := make([][]string, 0, len(mp))
    for _, value := range mp {
        result = append(result, value)
    }

    return result
}
