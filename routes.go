package CVPortfolio

import "net/http"

//	routes
//
// initialises all the routes.
func routes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
}
