function findContentChildren(g: number[], s: number[]): number {
    let c: number = 0;

    g.sort((a, b) => a - b);
    s.sort((a, b) => a - b);

    for (let i = 0; i < s.length && c < g.length; ++i){
        if (g[c] <= s[i]){
            ++c
        }
    }

    return c
};
