impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut answer = vec![1; nums.len()];
        let mut l = 0;
        let mut r = nums.len() - 1;

        let (mut lv, mut rv) = (1, 1);

        loop {
            answer[l] = answer[l] * lv;
            answer[r] = answer[r] * rv;

            rv = rv * nums[r];
            lv = lv * nums[l];

            if r == 0 {
                break;
            }
            
            l += 1;
            r -= 1;
        }

        answer
    }
}
