function search(nums: number[], target: number): number {
    let [l, r] = [0, nums.length - 1];

    while (l <= r) {
        const m = ((l + r) / 2) | 0;

        if (nums[m] === target) {
            return m;
        } else if (nums[m] < target) {
            l = m + 1;
        } else {
            r = m - 1;
        }
    }

    return -1;
};
