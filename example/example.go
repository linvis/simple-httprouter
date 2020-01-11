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

func helloParam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello Param")
}

func main() {
	r := router.New()

	r.GET("/get", helloGet)

	r.POST("/post", helloPost)

	r.GET("/user/:name", helloParam)

	r.Run(":5000")
}
