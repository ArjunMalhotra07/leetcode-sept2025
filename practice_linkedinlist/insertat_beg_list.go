package practice_linkedinlist

type ListNode struct {
	Num  int
	Next *ListNode
}

func InsertAtBeg(Head *ListNode, newNode *ListNode) *ListNode {
	if Head == nil {
		return newNode
	}
	newNode.Next = Head
	return newNode
}

func InsertAtLast(Head *ListNode, newNode *ListNode) *ListNode {
	if Head == nil {
		return newNode
	}
	Head.Next = InsertAtLast(Head.Next, newNode)
	return Head
}
