package main

import (
	"net/http"
	"github.com/tssaini/golang-webserver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}