package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	HumanInternalError = "Что-то пошло не так"

	ErrAlreadyExist    = "already exists"
	ErrAlreadyExistRus = "уже существует"
)

//handlers constant variables
var (
	SessionCookieName           = "session_token"
	SessionCookieExpirationTime = int64(120) //in seconds
)

//ErrResponse api error response common structure
type ErrResponse struct {
	Errors      []string `json:"errors"`
	HumanErrors []string `json:"human_errors"`
}

func respond500(w http.ResponseWriter) {
	errResp := ErrResponse{
		Errors:      []string{http.StatusText(http.StatusInternalServerError)},
		HumanErrors: []string{HumanInternalError},
	}
	resp, _ := json.Marshal(errResp)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(resp))
}

func errorToResponse(errtext string, humanerrtext string) []byte {
	errResp := ErrResponse{
		Errors:      []string{errtext},
		HumanErrors: []string{humanerrtext},
	}
	resp, _ := json.Marshal(errResp)
	return resp
}
