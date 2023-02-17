class Solution {
    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function rearrangeArray($nums) {
        $n = count($nums);
        $answ = array_fill(0, $n, 0);

        $positive = 0;
        $negative = 1;

        for ($i = 0; $i < $n; ++$i) {
            if ($nums[$i] > 0) {
                $answ[$positive] = $nums[$i];
                $positive += 2;
            } else {
                $answ[$negative] = $nums[$i];
                $negative += 2;
            }
        }

        return $answ;
    }
}
