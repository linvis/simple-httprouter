// version 1: name matching
// versino 2: parameter matching
package router

import "net/http"

import "fmt"

//Router struct
type Router struct {
	urlTree *Node
}

//HandlerFunc used for http GET/POST/... callback
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//New new a httprouter
func New() *Router {
	router := &Router{}

	router.urlTree = InitRouter()

	return router
}

//GET add get method
func (router *Router) GET(url string, handlers ...HandlerFunc) {
	router.urlTree.AddURL(url, handlers)
}

//POST add post method
func (router *Router) POST(url string, handlers ...HandlerFunc) {
	router.urlTree.AddURL(url, handlers)
}

//PUT method
func (router *Router) PUT(url string, handlers ...HandlerFunc) {
	router.urlTree.AddURL(url, handlers)
}

//DELETE method
func (router *Router) DELETE(url string, handlers ...HandlerFunc) {
	router.urlTree.AddURL(url, handlers)
}

//ServeHTTP override ServeHTTP
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("server url:", r.URL.Path)

	node := router.urlTree.Search(r.URL.Path)
	if node != nil {
		for _, f := range node.handlers {
			f(w, r)
		}
	}
}

//Run run httprouter
func (router *Router) Run(path string) {
	http.ListenAndServe(path, router)
}
