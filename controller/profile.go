package controller

import "net/http"

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
}

func getProfile() {

}
