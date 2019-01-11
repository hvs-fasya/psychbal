package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/chusha/internal/api/handlers/ws"
	"github.com/hvs-fasya/chusha/internal/engine"
	"github.com/hvs-fasya/chusha/internal/models"
)

//TabsGet get tabs list
func TabsGet(w http.ResponseWriter, r *http.Request) {
	var enabled bool
	enabled, err := strconv.ParseBool(r.URL.Query().Get("enabled"))
	if err != nil {
		enabled = true
	}
	tabs, err := engine.DB.TabsGet(enabled)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("[DB] TabsGet error %v", err)
		errResp := ErrResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		}
		resp, _ := json.Marshal(errResp)
		w.Write([]byte(resp))
		return
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.Marshal(tabs)
	w.Write(resp)
}

//TabsSet set tabs state
func TabsSet(w http.ResponseWriter, r *http.Request) {
	var e error
	body, e := ioutil.ReadAll(r.Body)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("[HTTP] failed to read TabsSet request body")
		errResp := ErrResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		}
		resp, _ := json.Marshal(errResp)
		w.Write([]byte(resp))
		return
	}

	var tabs = []*models.Tab{}
	e = json.Unmarshal(body, &tabs)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("[HTTP] failed to unmarshal TabsSet request body: %s", string(body))
		errResp := ErrResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		}
		resp, _ := json.Marshal(errResp)
		w.Write([]byte(resp))
		return
	}
	e = engine.DB.TabsSet(tabs)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("[DB] TabsSet error: %s", e)
		errResp := ErrResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		}
		resp, _ := json.Marshal(errResp)
		w.Write([]byte(resp))
		return
	}
	newTabs, e := engine.DB.TabsGet(false)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("[DB] TabsGet error %v", e)
		errResp := ErrResponse{
			Errors: []string{http.StatusText(http.StatusInternalServerError)},
		}
		resp, _ := json.Marshal(errResp)
		w.Write([]byte(resp))
		return
	}
	msg := &ws.Message{
		Msg:     "refresh tabs",
		Payload: newTabs,
	}
	ws.OutChann <- msg
	w.WriteHeader(http.StatusOK)
}
