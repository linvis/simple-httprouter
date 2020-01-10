// version 1: name matching
// versino 2: parameter matching
package router

import "net/http"

import "fmt"

//Router struct
type Router struct {
	urlPool map[string]*Node
}

//HandlerFunc used for http GET/POST/... callback
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//New new a httprouter
func New() *Router {
	router := &Router{}

	router.urlPool = make(map[string]*Node)

	return router
}

//GET add get method
func (router *Router) GET(url string, handlers ...HandlerFunc) {
	root, ok := router.urlPool["GET"]
	if !ok {
		root = InitNode()
		router.urlPool["GET"] = root
	}
	root.AddURL(url, handlers)
}

//POST add post method
func (router *Router) POST(url string, handlers ...HandlerFunc) {

	root, ok := router.urlPool["POST"]
	if !ok {
		root = InitNode()
		router.urlPool["POST"] = root
	}
	root.AddURL(url, handlers)
}

//PUT method
func (router *Router) PUT(url string, handlers ...HandlerFunc) {

	root, ok := router.urlPool["PUT"]
	if !ok {
		root = InitNode()
		router.urlPool["PUT"] = root
	}
	root.AddURL(url, handlers)
}

//DELETE method
func (router *Router) DELETE(url string, handlers ...HandlerFunc) {

	root, ok := router.urlPool["DELETE"]
	if !ok {
		root = InitNode()
		router.urlPool["DELETE"] = root
	}
	root.AddURL(url, handlers)
}

//ServeHTTP override ServeHTTP
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server method,  url:", r.Method, r.URL.Path)

	root, ok := router.urlPool[r.Method]
	if !ok {
		return
	}

	node := root.Search(r.URL.Path)
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
