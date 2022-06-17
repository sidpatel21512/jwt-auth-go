package controller

import "net/http"

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	success := getHealth()

	w.Write([]byte(success))
}

func getHealth() string {
	return "<h2 style='text-align:center'>Success</h2>"
}
