impl Solution {
    pub fn remove_elements(mut head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
        let mut dummy = None;
        let mut tail = &mut dummy;

        while let Some(mut node) = head.take() {
            head = node.next.take();

            if node.val != val {
                *tail = Some(node);
                tail = &mut tail.as_mut().unwrap().next;
            }
        }

        dummy
    }
}
