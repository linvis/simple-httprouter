package router

import "testing"

func TestRouter(t *testing.T) {
	root := InitRouter()

	str1 := "/user/info"
	root.AddURL(str1)
	if root.Search(str1) == false {
		t.Error("router not matching")
	}

	str2 := "/user/action"
	root.AddURL(str2)
	if root.Search(str2) == false {
		t.Error("router not matching")
	}

	str3 := "/user/actio"
	if root.Search(str3) == true {
		t.Error("router not matching")
	}
}
