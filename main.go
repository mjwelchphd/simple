package main

import (
	"fmt"
  "net"
	"net/http"
  "os"
)

type Application struct {
	mux *http.ServeMux
}

func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  host, port, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Printf("Hello, this is app router\n\n%+v\n\napp route host: '%s', port: %s\n\n", r, host, port)
	a.mux.ServeHTTP(w, r)
}

func crap(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is crap\n")
}

func welcome(w http.ResponseWriter, r *http.Request) {
  switch r.RequestURI {
  case "/":
    fmt.Fprintf(w, "This is the welcome page\n\n")
  default:
    fmt.Fprintf(w, "This is an attempted hack\n\n")
  }
}

func main() {
	app := &Application{ mux: http.NewServeMux() }
	app.mux.HandleFunc("/crap", crap)
  app.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  app.mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/images"))))
  app.mux.HandleFunc("/", welcome)
	http.ListenAndServe(":9292", app)
}
