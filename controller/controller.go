package controller

import (
	"fmt"
	"net/http"
)

// Home defines the home controller and registers the route
type Home struct {
}

// RegisterRoutes ...
func (h *Home) RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("this is the data API V 1.0"))
	})

	http.HandleFunc("/home", h.homeHandle)
}

func (h *Home) homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home route")
}
