package server

import (
	"fmt"
	"net/http"
)

// Main function to run server
func Main() error {
	fs := http.FileServer(http.Dir("react-demo/build"))
	http.Handle("/", fs)

	fmt.Println("Server react-demo at http://localhost:3000")
	return http.ListenAndServe(":3000", nil)
}
