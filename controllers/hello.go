package controllers

import "net/http"

var HelloHandler = func(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World"))
}
