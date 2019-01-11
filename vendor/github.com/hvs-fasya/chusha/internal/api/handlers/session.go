package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"

	"github.com/hvs-fasya/chusha/internal/engine"
	"github.com/hvs-fasya/chusha/internal/redis-client"
)

//SessionCreate login user
func SessionCreate(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte{})
		return
	}
	loginData := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	err = json.Unmarshal(payload, &loginData)
	user, err := engine.DB.UserCheck(loginData.Login, loginData.Password)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	sessionToken := uuid.NewV4().String()
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(SessionCookieExpirationTime) * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})
	err = redis_client.RedisClient.Client.Set(sessionToken, user.Nickname, time.Duration(SessionCookieExpirationTime)*time.Minute).Err()
	if err != nil {
		log.Error().Msgf("redis SET SESSION TOKEN request error: %s", err)
		respond500(w)
		return
	}
	resp, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//SessionDestroy logout user
func SessionDestroy(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(SessionCookieName)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	sessionToken := strings.Replace(c.Value, SessionCookieName+"=", "", 1)
	valid, _ := redis_client.RedisClient.Client.Exists(sessionToken).Result()
	if valid == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, err = redis_client.RedisClient.Client.Del(sessionToken).Result()
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sessionToken))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sessionToken))
}

//SessionCheck check session is valid
func SessionCheck(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(SessionCookieName)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	sessionToken := strings.Replace(c.Value, SessionCookieName+"=", "", 1)
	username, err := redis_client.RedisClient.Client.Get(sessionToken).Result()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := engine.DB.UserGetByName(username)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			log.Error().Msgf("database USER GET BY NAME request error: %s", err)
			respond500(w)
			return
		}
	}
	resp, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
