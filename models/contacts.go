package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {

	ContactName string `json:"contact_name"`
	PhoneNumber string `json:"phone_number"`
	UserId int64 `json:"user_id"`
	gorm.Model
}

func (con *Contact) Validate() (map[string]interface{}, bool) {

	resp := make(map[string] interface{})

	if con.ContactName == "" {
		resp["status"] = false
		resp["message"] = "Contact name is required."
		return resp, false
	}

	if con.PhoneNumber == "" {
		resp["status"] = false
		resp["message"] = "Phone number is required."
		return resp, false
	}

	if con.UserId == 0 {
		resp["status"] = false
		resp["message"] = "User is required."
		return resp, false
	}

	user := &User{}
	GetConn().Table("users").Where("id = ?", con.UserId).First(user)
	if user.Username == "" {
		resp["status"] = false
		resp["message"] = "User not found."
		return resp, false
	}

	resp["status"] = true
	resp["message"] = "All required fields(s) present"
	return resp, true
}

func (contact *Contact) Create() (map[string]interface{}) {

	resp := make(map[string]interface{})
	con := &Contact{}
	GetConn().Where("phone_number = ? AND user_id = ?", contact.PhoneNumber, contact.UserId).First(con)
	if con.PhoneNumber != "" {
		resp["status"] = false
		resp["message"] = "This contact already exists in user's contact list"
		return resp;
	}

	GetConn().Create(contact)
	resp["status"] = true
	resp["message"] = "Contact has been created."
	resp["contact"] = GetContact(contact.ID)
	return resp
}
