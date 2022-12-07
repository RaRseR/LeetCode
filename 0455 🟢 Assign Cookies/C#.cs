public int FindContentChildren(int[] g, int[] s) {
  int c = 0;

  Array.Sort(g);
  Array.Sort(s);

  for (int i = 0; i < s.Length && c < g.Length; ++i) {
    if (g[c] <= s[i]) {
      ++c;
    }
  }

  return c;
  }
