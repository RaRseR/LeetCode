use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();

        for (i, n) in nums.iter().enumerate() {
            if let Some(p) = map.get(n) {
                return vec![*p as i32, i as i32];
            }

            map.insert(target - n, i as i32);
        }
        
        unreachable!("Result should have been found");
    }
}
