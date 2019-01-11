package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"

	"github.com/hvs-fasya/chusha/internal/engine"
	"github.com/hvs-fasya/chusha/internal/models"
	"github.com/hvs-fasya/chusha/internal/redis-client"
	"github.com/hvs-fasya/chusha/internal/utils"
)

//UserRegister register user
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var err error
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Msgf("http USER CREATE read body error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorToResponse("failed to read request body", HumanInternalError)
		return
	}
	user := new(models.UserNewInput)
	err = json.Unmarshal(payload, &user)
	if err != nil {
		log.Error().Msgf("http USER CREATE unmarshal body error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		errorToResponse("failed to unmarshal request body", HumanInternalError)
		return
	}
	errs := user.Validate()
	if len(errs) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		errResp := ErrResponse{
			Errors:      []string{},
			HumanErrors: []string{},
		}
		for _, e := range errs {
			errResp.Errors = append(errResp.Errors, e)
			errResp.HumanErrors = append(errResp.HumanErrors, e)
		}
		resp, _ := json.Marshal(errResp)
		w.Write(resp)
		return
	}
	err = engine.DB.UserCreate(user, utils.ClientRoleName)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23505" { //unique key login or email violation
				w.WriteHeader(http.StatusBadRequest)
				w.Write(errorToResponse("login or email "+ErrAlreadyExist, "логин или email "+ErrAlreadyExistRus))
				return
			}
			log.Error().Msgf("database USER CREATE request error: %s", err)
			respond500(w)
			return
		}
		log.Error().Msgf("database USER CREATE request error: %s", err)
		respond500(w)
		return
	}
	sessionToken := uuid.NewV4().String()
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionToken,
		Expires:  time.Now().Add(time.Duration(SessionCookieExpirationTime) * time.Second),
		Secure:   true,
		HttpOnly: true,
	})
	err = redis_client.RedisClient.Client.Set(sessionToken, user.Nickname, time.Duration(SessionCookieExpirationTime)*time.Minute).Err()
	if err != nil {
		log.Error().Msgf("redis SET SESSION TOKEN request error: %s", err)
		respond500(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}
