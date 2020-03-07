
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var startNode *ListNode
	var lastOfSortedNode *ListNode
	lastOfSortedNode = nil
	if l1 == nil {
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val <= l2.Val{
		startNode =l1
	} else {
		startNode = l2
	}
	for {
		if l1.Val <= l2.Val{
			if lastOfSortedNode !=nil {lastOfSortedNode.Next = l1}
			if l1.Next==nil {
				l1.Next = l2
				break
			}
			l1, l2, lastOfSortedNode= newTwoLists(l1,l2)
		} else {
			if lastOfSortedNode !=nil {lastOfSortedNode.Next = l2}
			if l2.Next == nil {
				l2.Next = l1
				break
			}
			l1, l2, lastOfSortedNode = newTwoLists(l2, l1)
		}
	}
	return startNode
}

func newTwoLists(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode , *ListNode) {
	var lastOfSortedNode *ListNode
	lastOfSortedNode = l1
	newl2 := l1.Next
	l1.Next  = l2
	newl1 := l2	
	return newl1, newl2, lastOfSortedNode
}