impl Solution {
    pub fn successful_pairs(mut spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
        potions.sort();

        for i in 0..spells.len() {
            let mut left = 0;
            let mut right = potions.len();

            while left < right {
                let mid = left + (right - left) / 2;

                if spells[i] as i64 * potions[mid] as i64 >= success {
                    right = mid;
                } else {
                    left = mid + 1;
                }
            }

            spells[i] = (potions.len() - left) as i32;
        }

        spells
    }
}
