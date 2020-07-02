package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/tssaini/golang-webserver/models"
	"regexp"
	"net/http"
)

type userController struct{
	userIDPattern *regexp.Regexp

}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/users"{
		switch r.Method {
			case http.MethodGet:
				uc.getAll(w, r)
			case http.MethodPost:
				uc.post(w,r)
			default:
				w.WriteHeader(http.StatusNotImplemented)
		}

	}else{
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0{
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil{
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
			case http.MethodGet:
				uc.get(id, w)
			default:
				w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request){
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter){
	u, err := models.GetUserByID(id)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request){
	u, err := uc.parseRequest(r)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error){
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil{
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *userController{
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}