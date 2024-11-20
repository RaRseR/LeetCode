impl Solution {
    pub fn take_characters(s: String, k: i32) -> i32 {
        let n: usize = s.len();

        if n < (k * 3) as usize {
            return -1;
        }

        let bytes: &[u8] = s.as_bytes();

        let mut count: Vec<i32> = vec![0; 3];

        for i in 0..n {
            count[(bytes[i] - b'a') as usize] += 1
        }

        if count[0] < k || count[1] < k || count[2] < k {
            return -1;
        }

        let mut result: i32 = i32::MAX;
        let mut left: usize = 0;

        for right in 0..n {
            count[(bytes[right] - b'a') as usize] -= 1;

            while count[0] < k || count[1] < k || count[2] < k {
                count[(bytes[left] - b'a') as usize] += 1;

                left += 1;
            }

            result = result.min((n - (right - left + 1)) as i32);
        }

        return result;
    }
}
