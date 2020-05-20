package server

import (
	"fmt"
	"html"
	"net/http"
)

var (
	_ http.Handler = (*handler)(nil)
)

type handler struct{}

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
