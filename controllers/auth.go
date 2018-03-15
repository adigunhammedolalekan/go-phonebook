package controllers

import (
	"net/http"
	"phonebook/models"
	"encoding/json"
	"fmt"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	if resp, ok := user.Validate(); !ok {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	if resp, ok := user.Create(models.GetConn()); !ok {
		Respond(w, resp)
	}else {
		Respond(w, resp)
	}
}

var Login = func(w http.ResponseWriter, r *http.Request) {

	resp := make(map[string]interface{})

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		resp["status"] = false
		resp["message"] = "Malformed payload"
		Respond(w, resp)
		return
	}

	resp = models.Login(user.Username, user.Password)
	Respond(w, resp)
}

func Respond(w http.ResponseWriter, data map[string]interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
