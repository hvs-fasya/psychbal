package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/hvs-fasya/chusha/internal/api/handlers"
	"github.com/hvs-fasya/chusha/internal/api/handlers/front"
	"github.com/hvs-fasya/chusha/internal/api/handlers/ws"
)

// Server wraps http server
type Server struct {
	httpServer *http.Server
}

// Run server
func (s *Server) Run(connstr string) {
	log.Info().Msg("Start server at " + connstr)
	e := http.ListenAndServeTLS(connstr, "cert.pem", "key.pem", NewRouter())
	if e != nil {
		log.Fatal().Err(e).Msg("Start server error")
	}
}

// NewRouter Создать - новый роутер
func NewRouter() *mux.Router {
	rt := new(mux.Router)
	// front
	rt.HandleFunc("/alive", handlers.Alive).Methods("GET")
	rt.HandleFunc("/", front.IndexHandler).Methods("GET")
	rt.HandleFunc(`/{file:favicon.+}`, front.FaviconHandler).Methods("GET")
	rt.HandleFunc(`/{js:.+\.js}`, front.JSHandler).Methods("GET")
	rt.HandleFunc("/fonts/{font}", front.FontsHandler).Methods("GET")
	// websocket
	rt.HandleFunc("/wss", ws.ServeWs)
	// api/v1
	apiRouter := rt.PathPrefix(("/api/v1")).Subrouter()
	apiRouter.Use(setApiHeaders)
	apiRouter.Use(respondOptions)

	apiRouter.HandleFunc("/tabs", handlers.TabsGet).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/tabs", handlers.TabsSet).Methods("PUT", "OPTIONS")
	apiRouter.HandleFunc("/session", handlers.SessionCheck).Methods("GET", "OPTIONS")
	apiRouter.HandleFunc("/session", handlers.SessionCreate).Methods("POST", "OPTIONS")
	apiRouter.HandleFunc("/session", handlers.SessionDestroy).Methods("DELETE", "OPTIONS")
	apiRouter.HandleFunc("/user", handlers.UserRegister).Methods("POST", "OPTIONS")
	return rt
}

func setApiHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "private, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "-1")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,Expires,If-Modified-Since,Keep-Alive,Origin,Pragma,UserDB-Agent,X-Requested-With,X-Initiator-UserDB")
		next.ServeHTTP(w, r)
	})
}

func respondOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte{})
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
