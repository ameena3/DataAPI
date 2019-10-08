package main

import (
	"net/http"

	"github.com/ameena3/DataAPI/controller"
)

func main() {
	var h controller.Home
	h.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
