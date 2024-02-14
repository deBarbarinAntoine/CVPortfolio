package CVPortfolio

import (
	"log"
	"net/http"
)

// rootHandler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// Index page handler.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	err := tmpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Fatal(err)
	}
}
