package practice_linkedinlist

func GetIntersectionNode(h1, h2 *ListNode) *ListNode {
	curr1 := h1
	curr2 := h2
	l1, l2 := 0, 0
	for curr1 != nil {
		l1 += 1
		curr1 = curr1.Next
	}
	for curr2 != nil {
		l2 += 1
		curr2 = curr2.Next
	}
	if l1 > l2 {
		return collisionPoint(l1-l2, h2, h1)
	} else {
		return collisionPoint(l2-l1, h1, h2)
	}
}
func collisionPoint(diff int, Smaller, Bigger *ListNode) *ListNode {
	for i := 0; i < diff; i++ {
		Bigger = Bigger.Next
	}
	for Smaller != Bigger {
		Smaller = Smaller.Next
		Bigger = Bigger.Next
	}
	return Smaller
}
