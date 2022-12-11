func containsDuplicate(nums []int) bool {
   hash := map[int]bool{}

   for _, n := range nums {
        if _, ok := hash[n]; ok {
           return true
        }
        hash[n] = true
   }

   return false 
}
