type CustomStack struct {
    stack []int
    size int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{
        stack: make([]int, maxSize),
        size: 0,
    }
}


func (this *CustomStack) Push(x int)  {
    if this.size < len(this.stack) {
        this.stack[this.size] = x
        this.size++
    }
}


func (this *CustomStack) Pop() int {
    if this.size != 0 {
        val := this.stack[this.size - 1]

        this.size--

        return val
    }

    return -1
}


func (this *CustomStack) Increment(k int, val int)  {
    for i := 0; i < min(k, this.size); i++ {
        this.stack[i] += val
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
