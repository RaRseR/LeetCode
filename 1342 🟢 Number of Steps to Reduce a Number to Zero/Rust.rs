impl Solution {
    pub fn number_of_steps(mut num: i32) -> i32 {
        let mut steps = 0;

        while num != 0 {
            if num & 1 == 0 {
                num >>= 1;
            } else {
                num -= 1;
            }

            steps += 1;
        }

        steps
    }
}
