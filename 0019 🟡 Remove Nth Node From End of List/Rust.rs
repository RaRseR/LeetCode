impl Solution {
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        helper(head, n).0
    }
}

fn helper(head: Option<Box<ListNode>>, n: i32) -> (Option<Box<ListNode>>, usize) {
    match head {
        Some(mut node) => {
            let (prev, num) = helper(node.next.take(), n);
            if n == num as i32 {
                (prev, num + 1)
            } else {
                node.next = prev;
                (Some(node), num + 1)
            }
        },
        None => (None, 1)
    }
}
