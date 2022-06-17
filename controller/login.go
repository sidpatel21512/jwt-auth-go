package controller

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
}

func login() {

}
