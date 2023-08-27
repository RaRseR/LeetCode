impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = Vec::with_capacity(s.len());

        for c in s.chars() {
            match c {
                '{' => stack.push('}'),
                '(' => stack.push(')'),
                '[' => stack.push(']'),
                '}'|')'|']' if Some(c) != stack.pop() => return false,
                _ => (),
            }
        }

        return stack.is_empty();
    }
}
