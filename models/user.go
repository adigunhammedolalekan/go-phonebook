package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"os"
	"fmt"
)


type Token struct {
	User uint
	Username string
	jwt.StandardClaims
}

type User struct {

	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string

}


func (u *User) Create(db *gorm.DB) (map[string] interface{}, bool) {

	tempUser := User{}
	resp := make(map[string] interface{})

	GetConn().Where("username = ?", u.Username).First(&tempUser)
	if tempUser.Username != "" {
		resp["status"] = false
		resp["message"] = "Username already exists"
		return resp, false
	}

	GetConn().Where("email = ?", u.Email).First(&tempUser)
	if tempUser.Email != "" {
		resp["status"] = false
		resp["message"] = "Email address already in use"
		return resp, false
	}

	encodePass, _  := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(encodePass)
	db.Save(u)

	u.Password = ""
	resp["status"] = true
	resp["message"] = "Account has been created."
	resp["account"] = u

	tokenSecret := os.Getenv("token_password")
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{User: u.ID, Username: u.Username})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		fmt.Print(err)
	}
	resp["token"] = tokenString
	return resp, true
}

func (u *User) Validate() (map[string]interface{}, bool) {

	response := make(map[string]interface{})

	if strings.TrimSpace(u.Username) == "" {
		response["status"] = false
		response["message"] = "Username cannot be empty"
		return response, false
	}

	if strings.TrimSpace(u.Email) == "" {
		response["status"] = false
		response["message"] = "Email cannot be empty"
		return response, false
	}

	if strings.TrimSpace(u.Password) == "" {
		response["status"] = false
		response["message"] = "Password cannot be empty"
		return response, false
	}

	response["status"] = true
	response["message"] = "All Required field(s) present."
	return response, true
}

func Login(username, password string) (map[string]interface{}) {

	resp := make(map[string]interface{})

	user := &User{}
	GetConn().Where("username = ?", username).First(user)
	if user.Username == "" {
		resp["status"] = false
		resp["message"] = "User does not exist."
		return resp
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		resp["status"] = false
		if err == bcrypt.ErrMismatchedHashAndPassword {
			resp["message"] = "Authentication failed. Password incorrect"
		}else {
			resp["message"] = "Authentication failed. Unknown error."
		}
		return resp
	}

	user.Password = ""
	resp["status"] = true
	resp["message"] = "Login success"
	resp["account"] = user

	tokenSecret := os.Getenv("token_password")
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{User:user.ID, Username: user.Username})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		fmt.Println(err)
	}

	resp["token"] = tokenString
	return resp
}