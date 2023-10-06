func reverseWords(s string) string {
    words := strings.Split(s, " ")

    for w := 0; w < len(words); w++ {
        word := []rune(words[w])

        for i, j := 0, len(word) - 1; i < j; i, j = i+1, j-1 {
		    word[i], word[j] = word[j], word[i]
	    }

        words[w] = string(word)
    }

    return strings.Join(words, " ")
}
