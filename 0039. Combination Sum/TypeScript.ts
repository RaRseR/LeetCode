function combinationSum(candidates: number[], target: number): number[][] {
    candidates.sort((a, b) => a - b);
    
    const result : number[][] = [];
    
    helper(candidates, target, [], result)
    
    return result
};

function helper(n : number[], t : number, p : number[], r : number[][]){
    if (t < 0){
        return
    }
    
    if (t === 0) {
        r.push(p)
        return
    }
    
    for (let i = 0; i < n.length; ++i){  
        helper(n.slice(i), t - n[i], [...p, n[i]], r)
    }
    
}
