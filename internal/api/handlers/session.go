package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/psychbal/internal/models"
)

//AuthHandler http handler for form.html file
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, StaticPath+"form.html")
}

//LoginHandler handle login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if e != nil {
		log.Error().Msgf("parse LOGIN form error: %s", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	loginData := new(models.LoginInput)
	e = decoder.Decode(loginData, r.PostForm)
	if e != nil {
		log.Error().Msgf("decode LOGIN form data error: %s", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var user *models.User
	//user, e := oauth.ADAuth.CheckCreds(loginData)
	//if e != nil {
	//	switch e {
	//	case oauth.UnauthorizedError:
	//		//todo: handle 403
	//		w.WriteHeader(http.StatusForbidden)
	//		return
	//	default:
	//		log.Error().Msgf("check AD AUTH error: %s", e)
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}

	setCookie(user, w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//LogoutHandler handle logout request
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    SessionCookieName,
		Value:   "",
		Expires: time.Now().Add(-24 * time.Hour),
		Path:    "/",
		//todo: разобраться будет ли кука secure если https только на нгинксе
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/auth", http.StatusSeeOther)
}

func setCookie(user *models.User, w http.ResponseWriter, r *http.Request) {
	claims := map[string]string{
		"id":   strconv.FormatInt(user.ID, 10),
		"name": user.Name,
	}
	encoded, err := Cookie.Encode(SessionCookieName, claims)
	if err == nil {
		cookie := &http.Cookie{
			Name:    SessionCookieName,
			Value:   encoded,
			Expires: time.Now().Add(time.Duration(SessionCookieExpirationTime) * time.Minute),
			Path:    "/",
			//todo: разобраться будет ли кука secure если https только на нгинксе
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	} else {
		log.Error().Msgf("ENCODE COOKIE error: %s", err)
	}
}
