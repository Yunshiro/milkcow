// Package milk_http is about http methods,
// and create the RESTful API
package milk_http

import (
	"fmt"
	"net/http"
)

type MilkFunc func(w http.ResponseWriter, r *http.Request)

// GET is a http method.
// pattern: the api url pattern, like '/user', '/book/page'
// port: the server's part number, like ':8080', ':7777'
// handler: the logic part, write one function that return a MilkFunc.
func GET(pattern string, port string, handler MilkFunc) {
	http.HandleFunc(pattern, handler)
	ListenErr := http.ListenAndServe(port, nil)
	if ListenErr != nil {
		fmt.Printf("ListenErr: %v\n", ListenErr)
	}
}
