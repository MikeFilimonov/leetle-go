package leethashmaps

type ListNode struct {
	Val  int
	Next *ListNode
}

func FloydCycleDetection(node *ListNode) bool {

	if node == nil || node.Next == nil {
		return false
	}
	slow := node
	fast := node.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false

}

func SlowLinkedListCycle(node *ListNode) bool {

	if node == nil {
		return false
	}
	buffer := make(map[*ListNode]bool)
	for node.Next != nil {

		if _, found := buffer[node]; found {
			return true
		}
		buffer[node] = true
		node = node.Next

	}

	return false

}
