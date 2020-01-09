package main

import (
	"fmt"
	"net/http"
	"router"
)

func helloPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello POST")
}
func helloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello GET")
}

func main() {
	r := router.New()

	r.GET("/post", helloPost)

	r.POST("/get", helloGet)

	r.Run(":5000")
}
