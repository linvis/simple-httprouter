// version 1: name matching
// versino 2: parameter matching
package router

type Node struct {
	Val      string
	IsEnd    bool
	Children []*Node
}

func InitRouter() *Node {
	return &Node{}
}

func (this *Node) FindNode(c rune) *Node {
	var ans *Node = nil

	for _, next := range this.Children {
		if string(c) == next.Val {
			ans = next
			break
		}
	}

	return ans
}

func (this *Node) AddURL(url string) {
	head := this

	if len(url) <= 0 {
		return
	}

	for _, c := range url {
		next := head.FindNode(c)
		if next == nil {
			next = &Node{}
			next.Val = string(c)
			head.Children = append(head.Children, next)
		}

		head = next
	}

	head.IsEnd = true
}

func (this *Node) Search(url string) bool {
	head := this

	if len(url) <= 0 {
		return false
	}

	for _, c := range url {
		next := head.FindNode(c)
		if next == nil {
			return false
		}

		head = next
	}

	if head.IsEnd == true {
		return true
	} else {
		return false
	}
}
