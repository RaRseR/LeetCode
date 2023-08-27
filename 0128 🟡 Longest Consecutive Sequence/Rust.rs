use std::collections::{HashMap, HashSet};

impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut answer: usize = 0;
        let set: HashSet<i32> = nums.into_iter().collect();

        for &num in &set {
            if !set.contains(&(num - 1)) {
                let count = (num..).take_while(|x| set.contains(x)).count();
                answer = answer.max(count);
            }
        }

        answer as i32
    }
}
