impl Solution {
    pub fn merge_alternately(w1: String, w2: String) -> String {
        let mut c1 = w1.chars();
        let mut c2 = w2.chars();

        let mut res = String::with_capacity(w1.len() + w2.len());

        for _ in 0..w1.len().min(w2.len()) {
            res.push(c1.next().unwrap());
            res.push(c2.next().unwrap());
        }

        res.extend(c1);
        res.extend(c2);

        res
    }
}
