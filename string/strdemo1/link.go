package strdemo1

import "fmt"

var (
	head *LinkNode
)

func GetHead()*LinkNode{
	return head
}

type LinkNode struct {
	value int
	next  *LinkNode
}

func init() {
	head = new(LinkNode)
}

func Add(head *LinkNode, val int) {
	if head.next == nil {
		head.next = &LinkNode{value: val, next: nil}
		return
	} else {
		p := head
		for p.next != nil {
			p = p.next
		}
		p.next = &LinkNode{value: val, next: nil}
	}
}

func TravelLink(head *LinkNode) {
	p := head.next
	for p != nil {
		fmt.Printf("%v ", p.value)
		p = p.next
	}
	fmt.Println()
}

