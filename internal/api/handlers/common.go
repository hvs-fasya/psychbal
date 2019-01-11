package handlers

import (
	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
)

var (
	SessionCookieName = "session_id"
	CtxUserKey        = "claims"

	decoder                     = schema.NewDecoder()
	SessionCookieExpirationTime = int64(240) //in minutes
	Cookie                      *securecookie.SecureCookie
)

//CookieCfg secure cookie config structure
type CookieCfg struct {
	HashKey  []byte
	BlockKey []byte
}

//InitCookieHandling initialize secure cookie handling
func InitCookieHandling(cfg *CookieCfg) {
	Cookie = securecookie.New(cfg.HashKey, cfg.BlockKey)
}
