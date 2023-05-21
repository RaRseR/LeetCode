use std::cmp::{max, min};

impl Solution {
    pub fn max_product(nums: Vec<i32>) -> i32 {
        let (mut result, mut cmax, mut cmin) = (nums[0], nums[0], nums[0]);
        
        for i in 1..nums.len() {
            if nums[i] < 0 {
                let c = cmax;
                cmax = cmin;
                cmin = c;
            }

            cmax = max(nums[i], cmax * nums[i]);
            cmin = min(nums[i], cmin * nums[i]);

            result = max(result, cmax);
        }

        result
    }
}
