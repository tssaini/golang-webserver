package controllers

import (
	"net/http"
)

// RegisterControllers handle all paths
func RegisterControllers(){
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}