package router

import "strings"

type Node struct {
	Val      string
	IsEnd    bool
	handlers []HandlerFunc
	Children []*Node
}

func InitNode() *Node {
	return &Node{}
}

//FindDelimiter split by delimiter
func FindDelimiter(s string) []string {
	return strings.Split(s, "/")
}

func (this *Node) FindNode(path string) *Node {
	var ans *Node = nil

	for _, next := range this.Children {
		if path == next.Val {
			ans = next
			return ans
		}
	}

	for _, next := range this.Children {
		if next.Val == "*" {
			ans = next
			return ans
		}
	}

	return ans
}

func checkParam(path string) bool {
	for _, c := range path {
		if c == ':' {
			return true
		}
	}

	return false
}

func (this *Node) AddURL(url string, handlers []HandlerFunc) {
	head := this

	if len(url) <= 0 {
		return
	}

	pathes := FindDelimiter(url)

	for _, path := range pathes {
		val := path
		if checkParam(path) == true {
			val = "*"
		}

		next := head.FindNode(path)
		if next == nil {
			next = &Node{}
			next.Val = val
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

	pathes := FindDelimiter(url)

	for _, path := range pathes {
		next := head.FindNode(path)
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
