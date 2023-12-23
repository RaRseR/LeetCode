func hammingWeight(num uint32) (result int) {
    for i := 0; i < 32; i++ {
        result += int(num & 1)
        num >>= 1
    }
    
    return
}
