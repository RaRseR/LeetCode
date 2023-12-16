impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let n = s.len();

        if n != t.len() {
            return false;
        }

        let mut chars: [i16; 26] = [0; 26];

        for i in 0..n {
            chars[s.as_bytes()[i] as usize - 97] += 1;
            chars[t.as_bytes()[i] as usize - 97] -= 1;
        }

        for i in chars.iter() {
            if *i != 0 {
                return false;
            }
        }

        return true;
    }
}
