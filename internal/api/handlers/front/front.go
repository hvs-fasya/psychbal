package front

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var (
	StaticPath = ""
)

func InitFront(staticPath string) {
	StaticPath = staticPath
}

//IndexHandler http handler for index.html file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg(StaticPath + "index.html")
	http.ServeFile(w, r, StaticPath+"index.html")
}

//FaviconHandler http handler for favicon.ico file
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["file"]
	w.Header().Set("Content-Type", "image/x-icon")
	http.ServeFile(w, r, StaticPath+f)
}

//SWHandler http handler for service-worker.js file
func SWHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	http.ServeFile(w, r, StaticPath+"service-worker.js")
}

//ManifestHandler http handler for manifest.js file
func ManifestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, StaticPath+"manifest.json")
}

//JSHandler http handler for js files
func JSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	vars := mux.Vars(r)
	file := vars["js"]
	http.ServeFile(w, r, StaticPath+file)
}

//FontsHandler http handler for font files
func FontsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["font"]
	http.ServeFile(w, r, StaticPath+"fonts/"+f)
}

//CSSHandler http handler for font files
func CSSHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["css"]
	http.ServeFile(w, r, StaticPath+"css/"+f)
}

//IMGHandler http handler for font files
func IMGHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["img"]
	http.ServeFile(w, r, StaticPath+"img/"+f)
}

//IconsHandler http handler for font files
func IconsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := vars["img"]
	http.ServeFile(w, r, StaticPath+"img/icons/"+f)
}
