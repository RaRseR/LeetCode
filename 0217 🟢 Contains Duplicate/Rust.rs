use std::collections::HashSet;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut set: HashSet<i32> = HashSet::with_capacity(nums.len());

        for i in nums.iter() {
            if set.contains(i) {
                return true;
            }

            set.insert(*i);
        }

        false
    }
}
