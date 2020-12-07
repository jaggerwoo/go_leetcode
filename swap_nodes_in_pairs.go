// Given a linked list, swap every two adjacent nodes and return its head.

// You may not modify the values in the list's nodes, only nodes itself may be changed.

//  

// Example:

// Given 1->2->3->4, you should return the list as 2->1->4->3.


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    a := head
    b := head.Next
    c := head.Next.Next

    b.Next = a
    a.Next = swapPairs(c)

    return b
}