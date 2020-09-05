package demo10

type ListNode struct {
	Val  int
	Next *ListNode
}

func Create(data []int, index int) (head *ListNode) {
	if index >= len(data) {
		return
	}

	head = &ListNode{data[index], nil}
	head.Next = Create(data, index+1)
	return
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	p1 := pHead1
	p2 := pHead2
	if p1 == nil || p2 == nil {
		return nil
	}
	s1 := []*ListNode{}
	s2 := []*ListNode{}
	for p1 != nil {
		s1 = append(s1, p1)
		p1 = p1.Next
	}
	for p2 != nil {
		s2 = append(s2, p2)
		p2 = p2.Next
	}
	i, j := len(s1)-1, len(s2)-1
	var res *ListNode
	for i >= 0 && j >= 0 {
		if s1[i] == s2[j] {
			res = s1[i]
		} else {
			break
		}
		i--
		j--
	}
	return res
}

//因为从头到尾不好比较，从尾到头可以，这时应该考虑到栈
