/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]int)
	for head != nil {
		_, ok := hash[head]
		if ok {
			return true
		}
		hash[head] = head.Val
		head = head.Next
	}
	return false
}