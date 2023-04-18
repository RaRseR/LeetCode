func mergeAlternately(w1 string, w2 string) string {
    var res []byte

    for i := 0; i < len(w1) || i < len(w2); i++ {
        if i < len(w1) {
            res = append(res, w1[i])
        }
        
        if i < len(w2) {
            res = append(res, w2[i])
        }
    }
    
    return string(res)
}
