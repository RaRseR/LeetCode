func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }

    count := [26]int{}

    for i := 0; i < len(magazine); i++ {
        count[magazine[i] - 'a']++
    }

    for i := 0; i < len(ransomNote); i++ {
        count[ransomNote[i] - 'a']--
        if count[ransomNote[i] - 'a'] < 0 {
            return false
        }
    }

    return true
}
