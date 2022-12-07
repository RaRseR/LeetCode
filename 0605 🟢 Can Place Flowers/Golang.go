func canPlaceFlowers(f []int, n int) bool {
    c := 0

    for i := range f {
        if (f[i] == 0) {
            l := i == 0 || f[i - 1] == 0;
            r := i == len(f) - 1 || f[i + 1] == 0;
                
            if l && r {
                f[i] = 1;
                c++;
            }
        }
    }

    return c >= n;
}
