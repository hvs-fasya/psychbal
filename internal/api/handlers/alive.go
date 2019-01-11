package handlers

import (
	"net/http"
)

// Alive - функция
func Alive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HTTP OK\n"))
}
