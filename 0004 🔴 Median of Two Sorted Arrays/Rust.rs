impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let (mut nums1, mut nums2) = if nums1.len() > nums2.len() {
            (nums2, nums1)
        } else {
            (nums1, nums2)
        };

        let n1 = nums1.len();
        let n2 = nums2.len();
        
        let mut left = 0;
        let mut right = n1;

        while left <= right {
            let partition1 = (left + right) / 2;
            let partition2 = (n1 + n2 + 1) / 2 - partition1;

            let max_left1 = if partition1 == 0 {
                i32::MIN
            } else {
                nums1[partition1 - 1]
            };
            let min_right1 = if partition1 == n1 {
                i32::MAX
            } else {
                nums1[partition1]
            };

            let max_left2 = if partition2 == 0 {
                i32::MIN
            } else {
                nums2[partition2 - 1]
            };
            let min_right2 = if partition2 == n2 {
                i32::MAX
            } else {
                nums2[partition2]
            };

            if max_left1 <= min_right2 && max_left2 <= min_right1 {
                if (n1 + n2) % 2 == 0 {
                    return (max_left1.max(max_left2) + min_right1.min(min_right2)) as f64 / 2.0;
                } else {
                    return max_left1.max(max_left2) as f64;
                }
            } else if max_left1 > min_right2 {
                right = partition1 - 1;
            } else {
                left = partition1 + 1;
            }
        }

        return 0.0;
    }
}
