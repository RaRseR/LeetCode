impl Solution {
    pub fn delete_middle(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut size = 0;
        let mut node = head.as_ref();

        while node.is_some() { 
            node = node.unwrap().next.as_ref(); 
            size += 1; 
        }

        if size <= 1 { 
            return Option::None; 
        }

        size /= 2;
        let mut node = head.as_mut();

        while size > 1 { 
            node = node.unwrap().next.as_mut(); 
            size -= 1; 
        }

        node.unwrap().next = node.as_mut().unwrap().next.as_mut().unwrap().next.take();

        return head;
    }
}
