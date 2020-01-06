package main

import (
	"fmt"
	"net/http"
	"router"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	r := router.New()

	r.Get("/hello", homePage)

	r.Run(":5000")
}
