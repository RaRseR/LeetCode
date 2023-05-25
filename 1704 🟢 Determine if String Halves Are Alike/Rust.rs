impl Solution {
    pub fn halves_are_alike(s: String) -> bool {
        let helper = |s: &str| s.chars().filter(|&c| "aeiouAEIOU".contains(c)).count();

        helper(&s[..(s.len() / 2)]) == helper(&s[(s.len() / 2)..])
    }
}
