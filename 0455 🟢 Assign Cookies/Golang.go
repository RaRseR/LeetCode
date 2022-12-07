func findContentChildren(g []int, s []int) (c int) {
    sort.Ints(g)
    sort.Ints(s)

    for i := 0; i < len(s) && c < len(g); i++ {
        if (g[c] <= s[i]) {
            c++;
        }
    }

    return c
}
