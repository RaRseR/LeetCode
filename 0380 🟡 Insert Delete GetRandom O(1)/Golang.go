type RandomizedSet struct {
    mp map[int]int
    arr []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        mp: make(map[int]int),
        arr: make([]int, 0),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.mp[val]; ok {
        return false
    }

    this.mp[val] = len(this.arr)
    this.arr = append(this.arr, val)

    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if index, ok := this.mp[val]; ok {
        this.arr[index] = this.arr[len(this.arr) - 1]
        this.mp[this.arr[index]] = this.mp[val]
        this.arr = this.arr[:len(this.arr) - 1]

        delete(this.mp, val)
        
        return true
    }

    return false
}


func (this *RandomizedSet) GetRandom() int {
    return this.arr[rand.Intn(len(this.arr))]
}
