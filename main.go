package main

import (
	"factorial/cmd/calculus"
	"io"
	"net/http"
)
func main() {
	http.HandleFunc("/v1/factorial", calculus.NewCalculusHandler)
	http.HandleFunc("/v1/health-check", HealthCheckHandler)
	http.ListenAndServe(":5000", nil)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}


