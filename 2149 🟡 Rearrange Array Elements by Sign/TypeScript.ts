function rearrangeArray(nums: number[]): number[] {
    const n = nums.length;
    const answ = new Array<number>(n);

    let positive = 0,
        negative = 1;

    for (let i = 0; i < n; ++i) {
        if (nums[i] > 0) {
            answ[positive] = nums[i];
            positive += 2
        } else {
            answ[negative] = nums[i];
            negative += 2
        }
    }

    return answ;
};
