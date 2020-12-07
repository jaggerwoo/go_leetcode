// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

// Example:

// Given this linked list: 1->2->3->4->5

// For k = 2, you should return: 2->1->4->3->5

// For k = 3, you should return: 3->2->1->4->5

// Note:

// Only constant extra memory is allowed.
// You may not alter the values in the list's nodes, only nodes itself may be changed.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    // if head == nil || head.Next == nil {
    //     return head
    // }

    // tmp := head.Next
    // //判断剩下的节点数量是否足够
    // for i := 0; i < k-2; i++ {
    //     tmp = head.Next
    //     if tmp == nil {
    //         return head;
    //     }
    // }

    // a := head
    // b := head.Next
    // c := head.Next.Next
    // for j := 0; j < k; j++ {
    //     b.Next = a
    //     a.Next = reverseKGroup(c, j)
    // }
    // return b
}