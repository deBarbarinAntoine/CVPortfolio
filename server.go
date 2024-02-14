package CVPortfolio

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

var tmpl *template.Template

var (
	_, b, _, _ = runtime.Caller(0)
	path       = filepath.Dir(b) + "/"
)

// Establishing the fileserver to make assets directory available for the client.
func fileServer() {
	fs := http.FileServer(http.Dir(path + "assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// runServer runs the server in a goroutine and open the browser at the correct URL and then loops waiting for the "stop" input to end all goroutines.
func runServer() {
	port := ":8030"
	http.ListenAndServe(port, nil)
}

// Run is the public function that executes all necessary functions to run the server and website.
func Run() {
	var err error
	tmpl, err = template.ParseFiles(path + "templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	routes()
	fileServer()
	runServer()
}
