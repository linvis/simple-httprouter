package router

import "strings"

type Node struct {
	Val      string
	IsEnd    bool
	handlers []HandlerFunc
	Children []*Node
}

const (
	PrefixNotExist = iota
	PrefixMatchBoth
	PrefixMatchLeft
	PrefixMatchRight
	PrefixMatchSub
	PrefixMatchParam
)

type Prefix struct {
	ExistState int
	Pos        int
	Len        int
}

func InitNode() *Node {
	return &Node{}
}

//FindDelimiter split by delimiter
func FindDelimiter(s string) []string {
	return strings.Split(s, "/")[1:]
}

func findLongestPrefix(a string, b string) int {
	var i int = 0
	for ; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}
	if i == 0 {
		//no common
		return -1
	} else {
		return i
	}
}

//findNode return node exist, prefix length
func (root *Node) findNode(path string) *Prefix {
	prefix := &Prefix{
		ExistState: PrefixNotExist,
		Pos:        -1,
		Len:        0,
	}

	for i, next := range root.Children {
		prefixLen := findLongestPrefix(path, next.Val)

		// no common
		if prefixLen < 0 {
			continue
		} else {
			prefix.Pos = i
			prefix.Len = prefixLen
			if prefixLen == len(path) && prefixLen == len(next.Val) {
				prefix.ExistState = PrefixMatchBoth
			} else if prefixLen == len(path) {
				prefix.ExistState = PrefixMatchLeft
			} else if prefixLen == len(next.Val) {
				prefix.ExistState = PrefixMatchRight
			} else {
				prefix.ExistState = PrefixMatchSub
			}
			return prefix
		}
	}

	for i, next := range root.Children {
		if next.Val == "*" {
			prefix.Pos = i
			prefix.ExistState = PrefixMatchParam
			return prefix
		}
	}

	return prefix
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
	var next *Node = nil

	if len(url) <= 0 {
		return
	}

	pathes := FindDelimiter(url)

	for _, path := range pathes {
		val := path
		if checkParam(path) == true {
			val = "*"
		}

		prefix := head.findNode(path)

		if prefix.ExistState == PrefixNotExist || prefix.ExistState == PrefixMatchParam {
			next = &Node{}
			next.Val = val
			head.Children = append(head.Children, next)
		} else if prefix.ExistState == PrefixMatchBoth {
			next = head.Children[prefix.Pos]
		} else {
			if prefix.ExistState == PrefixMatchRight {
				next = head.Children[prefix.Pos]
			} else {
				next = &Node{}
				next.Val = path[:prefix.Len]
			}

			if prefix.ExistState != PrefixMatchRight {
				head.Children[prefix.Pos].Val = head.Children[prefix.Pos].Val[prefix.Len:]
				next.Children = append(next.Children, head.Children[prefix.Pos])
			}

			head.Children = append(head.Children[:prefix.Pos], head.Children[prefix.Pos+1:]...)
			head.Children = append(head.Children, next)

			if prefix.ExistState != PrefixMatchLeft {
				next.Children = append(next.Children, &Node{
					Val: path[prefix.Len:],
				})
				next = next.Children[len(next.Children)-1]
			}
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

	for i := 0; i < len(pathes); {
		path := pathes[i]

	LOOP:
		// idx, prefix := head.findNode(path)
		prefix := head.findNode(path)
		if prefix.ExistState == PrefixNotExist {
			return nil
		} else if prefix.ExistState == PrefixMatchLeft {
			return nil
		} else if prefix.ExistState == PrefixMatchRight {
			path = path[prefix.Len:]
			head = head.Children[prefix.Pos]
			goto LOOP
		} else if prefix.ExistState == PrefixMatchBoth || prefix.ExistState == PrefixMatchParam {
			head = head.Children[prefix.Pos]
			i++
		} else {
			path = path[prefix.Len:]
			goto LOOP
		}
	}

	if head.IsEnd == true {
		return head
	} else {
		return nil
	}
}
