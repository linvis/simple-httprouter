package router

import "testing"

func TestRouter(t *testing.T) {
	root := InitNode()

	str1 := "/user/info"
	root.AddURL(str1, nil)
	if root.Search(str1) == nil {
		t.Error("router not matching")
	}

	str2 := "/user/action"
	root.AddURL(str2, nil)
	if root.Search(str2) == nil {
		t.Error("router not matching")
	}

	str3 := "/user/actio"
	if root.Search(str3) != nil {
		t.Error("router not matching")
	}
}
