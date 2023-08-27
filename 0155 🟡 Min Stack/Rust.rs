struct MinStack {
    stack: Vec<(i32, i32)>,
}

impl MinStack {
    fn new() -> Self {
        MinStack { stack: vec![] }
    }
    
    fn push(&mut self, val: i32) {
        let min = val
            .min(
                *self
                .stack
                .last()
                .map(|(value, min)| min)
                .unwrap_or(&val)
            );
        self.stack.push((val, min));
    }
    
    fn pop(&mut self) {
        self.stack.pop();
    }
    
    fn top(&mut self) -> i32 {
        self.stack[self.stack.len() - 1].0
   }
    
    fn get_min(&self) -> i32 {
        self.stack[self.stack.len() - 1].1
    }
}
