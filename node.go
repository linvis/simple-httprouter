package router

type Node struct {
	Val      string
	IsEnd    bool
	handlers []HandlerFunc
	Children []*Node
}

func InitNode() *Node {
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

func (this *Node) AddURL(url string, handlers []HandlerFunc) {
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
	head.handlers = handlers
}

func (this *Node) Search(url string) *Node {
	head := this

	if len(url) <= 0 {
		return nil
	}

	for _, c := range url {
		next := head.FindNode(c)
		if next == nil {
			return nil
		}

		head = next
	}

	if head.IsEnd == true {
		return head
	} else {
		return nil
	}
}
