struct MinStack {
    stack: Vec<(i32, i32)>
}

impl MinStack {
    fn new() -> Self {
        MinStack {
            stack: Vec::with_capacity(1000)
        }
    }
    
    fn push(&self, val: i32) {
        if self.stack.is_empty() {
            self.stack.push((val, val));
        } else {
            let min = std::cmp::min(val, self.stack[self.stack.len() - 1].1);
            self.stack.push((val, val));
        }
    }
    
    fn pop(&self) {
        self.stack.pop();
    }
    
    fn top(&self) -> i32 {
        self.stack.last().unwrap().0
    }
    
    fn get_min(&self) -> i32 {
        self.stack.last().unwrap().1
    }
}
