package controllers

import (
	"net/http"
	"phonebook/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	contact := &models.Contact{}
	json.NewDecoder(r.Body).Decode(contact)

	if resp, ok := contact.Validate(); !ok {
		Respond(w, resp)
	}else {
		resp = contact.Create()
		Respond(w, resp)
	}
}

var GetUserContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")
	fmt.Printf("%s %d %s", r.URL.Path, user, "\n")
	params := mux.Vars(r)

	contacts := models.GetUserContact(params["id"])
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contacts" : contacts}
	Respond(w, resp)
}

var MyContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")
	contacts := models.GetUserContact(user)
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contacts" : contacts}
	Respond(w, resp)
}
