var phoneMap = [8]string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (result []string) {
    if len(digits) == 0 {
        return
    }

    helper("", digits, &result)

    return
}

func helper(combination, nextDigit string, result *[]string) {
    if len(nextDigit) == 0 {
        *result = append(*result, combination)

        return
    }

    chars := phoneMap[nextDigit[0] - '2']

    for _, char := range chars {
        helper(combination + string(char), nextDigit[1:], result)
    }
}
